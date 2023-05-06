package main

import (
	"flag"
	"log"
)

func main() {
	//token := mustToken()
	//tgClient := telegram.New(token)
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
