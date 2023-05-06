package telegram

import "bot-storage/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
}
