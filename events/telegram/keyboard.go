package telegram

import (
	"bot-storage/events"
	"bot-storage/lib/e"
	"encoding/json"
)

func createKeyBoard(nameButton ...string) (string, error) {
	var listButton []events.Button
	for _, v := range nameButton {
		listButton = append(listButton, events.Button{Text: v})
	}

	keyboard := events.KeyBoard{
		Buttons: [][]events.Button{listButton},
	}

	jsonKeyboard, err := json.Marshal(keyboard)
	if err != nil {
		return "", e.Wrap("json data encoding error", err)
	}

	return string(jsonKeyboard), nil
}
