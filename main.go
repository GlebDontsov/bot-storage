package main

import (
	"bot-storage/clients/telegram"
	"flag"
	"log"
)

const tgBotHost = "api.telegram.org"

func main() {
	tgClient := telegram.New(tgBotHost, mustToken())
	//fetcher := fether.New()
	//processor := processor.New(tgClient)
	//consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()
	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
