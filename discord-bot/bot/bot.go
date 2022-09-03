package bot

import (
	"discord-bot/config"
	"log"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Run() {
	// create bot session
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal(err)
		return
	}
	// make the bot a user
	user, err := goBot.User("@me")
	if err != nil {
		log.Fatal(err)
		return
	}
	BotID = user.ID
	goBot.AddHandler(messageHandler)
	err = goBot.Open()
	if err != nil {
		return
	}
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == BotID {
		return
	}
	/*if m.Content == "/환산 주스탯 사이트 검색" || m.Content == "3" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "사본 만들어야 사용이 가능해요!\n https://docs.google.com/spreadsheets/d/1hD-P07knD8FlzWvUOGm_3f-NMZxTTQ14IivqddUMXOY/copy")
	} else {
		_, _ = s.ChannelMessageSend(m.ChannelID, "오류")
	}*/
	// 도움 명령어
	/*switch m.Content {
	case "도움":
		_, _ = s.ChannelMessageSend(m.ChannelID,
			"명령어 사용 ex) .음악 기능사용 or .1 \n 1. 음악 기능사용\n 2. 메이플 전적검색\n 3. 환산 주스탯 사이트 검색\n 4.각종 주간보스 결정석 가격확인 \n 5. 주간보스 확인\n 6.일일퀘스트 확인\n 7.몬스터파크 확인\n 8.우르스 확인\n 9.코디 시뮬레이터\n 10.농장 검색\n 11. 메소시세")

	case "음악 기능 사용":

	case "3", "환산 주스탯 사이트 검색":
		_, _ = s.ChannelMessageSend(m.ChannelID, "사본 만들어야 사용이 가능해요!\n https://docs.google.com/spreadsheets/d/1hD-P07knD8FlzWvUOGm_3f-NMZxTTQ14IivqddUMXOY/copy")
	case "4", "각종 주간보스 결정석 가격확인":
		_, _ = s.ChannelMessageSend(m.ChannelID, "이번주 가격\n카오스 자쿰 9,741,285 \n하드 매그너스 11,579,023 \n하드 힐라 6,936,489 \n카오스 파풀라 26,725,937 \n피에르9,838,932 \n반반9,818,154 \n퀸9,806,780 \n카오스 핑크빈 7,923,110 \n시그너스 5,493,394 / 9,039,130 \n스우 33,942,566 / 118,294,192 \n데미안 35,517,853 / 112,480,613 \n가엔슬 46,935,874 / 155,492,141 \n루시드 48,058,319 / 57,502,626 / 131,095,655 \n윌 52,139,127 / 66,311,463 / 145,038,483 \n더스크 71,054,562 / 160,173,752 \n진 힐라 148,112,376 / 190,159,452 \n듄켈 76,601,412 / 168,609,280 \n(월간)검마 1,418,809,857 / 5,675,239,428 \n세렌 196,904,752 / 267,825,621 / 1,071,302,484 \n칼로스 300,000,000")

	//case "5"
	default:
		break
	}*/

}
