// https://github.com/chromedp/examples/blob/master/download_file/main.go
// https://github.com/chromedp/examples/blob/master/click/main.go
package mapleprofile

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/chromedp/cdproto/browser"
	"github.com/chromedp/chromedp"
)

func DownloadMapleProfile(ign string) {
	// create context
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	done := make(chan string, 1)
	chromedp.ListenTarget(ctx, func(v interface{}) {
		if ev, ok := v.(*browser.EventDownloadProgress); ok {
			completed := "(unknown)"
			if ev.TotalBytes != 0 {
				completed = fmt.Sprintf("%0.2f%%", ev.ReceivedBytes/ev.TotalBytes*100.0)
			}
			log.Printf("state: %s, completed: %s\n", ev.State.String(), completed)
			if ev.State == browser.DownloadProgressStateCompleted {
				done <- ev.GUID
				close(done)
			}
		}
	})

	// get working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Navigate(`https://maple.gg/u/`+ign),
		chromedp.WaitVisible(`button[data-target="#exampleModal"]`),
		chromedp.Click(`[data-target="#exampleModal"]`, chromedp.NodeReady),
		chromedp.WaitVisible(`#btn-save`),
		browser.SetDownloadBehavior(browser.SetDownloadBehaviorBehaviorAllowAndName).
			WithDownloadPath(wd).
			WithEventsEnabled(true),
		chromedp.Click(`#btn-save`, chromedp.NodeVisible),
	); err != nil && !strings.Contains(err.Error(), "net::ERR_ABORTED") {
		log.Fatal(err)
	}

	guid := <-done

	log.Printf("wrote %s", filepath.Join(wd, guid))
	os.Rename(filepath.Join(wd, guid), filepath.Join(wd, ign+".png"))
}
