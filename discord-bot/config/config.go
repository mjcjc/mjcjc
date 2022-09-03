package config

import (
	"encoding/json"
	"log"
	"os"
)

var (
	Token          string
	Prefix         string
	config         *configStruct
	RemoveCommands bool
)

type configStruct struct {
	Token          string `json:"token"`
	Prefix         string `json:"prefix"`
	RemoveCommands bool   `json:"removeCommands"`
}

func ReadConfig() error {
	file, err := os.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
		return err
	}
	Token = config.Token
	Prefix = config.Prefix
	RemoveCommands = config.RemoveCommands

	return nil
}
