package telegram

import (
	"bot-storage/lib/e"
	"encoding/json"
)

type KeyBoard struct {
	Buttons [][]Button `json:"keyboard"`
}

type Button struct {
	Text string `json:"text"`
}

func createKeyBoard(nameButton ...string) (string, error) {
	var listButton []Button
	for _, v := range nameButton {
		listButton = append(listButton, Button{Text: v})
	}

	keyboard := KeyBoard{
		Buttons: [][]Button{listButton},
	}

	jsonKeyboard, err := json.Marshal(keyboard)
	if err != nil {
		return "", e.Wrap("json data encoding error", err)
	}

	return string(jsonKeyboard), nil
}
