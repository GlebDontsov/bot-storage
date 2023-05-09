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

func ButtonMenu() (string, error) {
	rowButtons1 := []Button{
		{Text: ButtonGetLink},
		{Text: buttonAll},
	}

	rowButtons2 := []Button{
		{Text: buttonCount},
		{Text: ButtonHelp},
	}

	keyboard := KeyBoard{
		Buttons: [][]Button{rowButtons1, rowButtons2},
	}

	jsonKeyboard, err := json.Marshal(keyboard)
	if err != nil {
		return "", e.Wrap("json data encoding error", err)
	}

	return string(jsonKeyboard), nil
}

func ButtonSendLink() (string, error) {
	rowButtons1 := []Button{
		{Text: buttonFirst},
		{Text: buttonLast},
	}

	rowButtons2 := []Button{
		{Text: buttonRnd},
	}

	keyboard := KeyBoard{
		Buttons: [][]Button{rowButtons1, rowButtons2},
	}

	jsonKeyboard, err := json.Marshal(keyboard)
	if err != nil {
		return "", e.Wrap("json data encoding error", err)
	}

	return string(jsonKeyboard), nil
}
