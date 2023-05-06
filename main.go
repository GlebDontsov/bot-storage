package main

import (
	"bot-storage/storage/sqlite"
	"context"
	"flag"
	"log"
)

const (
	tgBotHost         = "api.telegram.org"
	sqliteStoragePath = "data/sqlite"
)

func main() {
	s, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatal("can not connect to storage: ", err)
	}

	if err := s.Init(context.TODO()); err != nil {
		log.Fatal("can not init storage: ", err)
	}

	//eventsProcessor := telegram.New(
	//	tgClient.New(tgBotHost, mustToken()),
	//	s,
	//)
	//tgClient := telegram.New(tgBotHost, mustToken())
	//fetcher := fetcher.New()
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
