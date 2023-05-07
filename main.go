package main

import (
	tgClient "bot-storage/clients/telegram"
	event_consumer "bot-storage/consumer/event-comsumer"
	"bot-storage/events/telegram"
	"bot-storage/storage/sqlite"
	"flag"
	"log"
)

const (
	tgBotHost         = "api.telegram.org"
	sqliteStoragePath = "data/sqlite/storage.db"
	batchSize         = 100
)

func main() {
	s, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatal("can not connect to storage: ", err)
	}

	if err := s.Init(); err != nil {
		log.Fatal("can not init storage: ", err)
	}

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		s,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()
	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
