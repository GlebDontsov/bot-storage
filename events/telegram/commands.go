package telegram

import (
	"bot-storage/lib/e"
	"bot-storage/storage"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s", text, username)

	if isAddCmd(text) {
		return p.savePage(chatID, text, username)
	}

	switch text {
	case AllCmd, buttonAll:
		return p.sendAll(chatID, username)
	case RndCmd, buttonRnd:
		return p.sendLink(chatID, username, RndCmd)
	case LastCmd, buttonLast:
		return p.sendLink(chatID, username, LastCmd)
	case FirstCmd, buttonFirst:
		return p.sendLink(chatID, username, FirstCmd)
	case HelpCmd, ButtonHelp:
		return p.sendHelp(chatID)
	case StartCmd:
		return p.sendHello(chatID)
	case ButtonGetLink:
		return p.choiceMethod(chatID)
	default:
		ID, err := strconv.Atoi(text)
		if err != nil {
			return p.sendTag(chatID, username, text)
		}
		return p.sendID(chatID, username, ID)
	}
}

func (p *Processor) savePage(chatID int, pageURL string, username string) error {
	page := &storage.Page{
		URL:      pageURL,
		UserName: username,
	}

	isExists, err := p.storage.IsExists(page)
	if err != nil {
		return e.Wrap("can not do command: save page", err)
	}

	if isExists {
		return p.tg.SendMessage(chatID, msgAlreadyExists, "")
	}

	if err := p.storage.Save(page); err != nil {
		return e.Wrap("can not do command: save page", err)
	}

	if err := p.tg.SendMessage(chatID, msgSaved, ""); err != nil {
		return e.Wrap("can not do command: save page", err)
	}

	return nil
}

func (p *Processor) sendLink(chatID int, username string, command string) (err error) {
	buttonsMenu, err := createKeyBoard(ButtonGetLink, buttonAll, ButtonHelp)
	if err != nil {
		return e.Wrap("can not do buttons", err)
	}

	var page *storage.Page

	switch command {
	case RndCmd:
		page, err = p.storage.PickRandom(username)
	case FirstCmd:
		page, err = p.storage.PickFirst(username)
	case LastCmd:
		page, err = p.storage.PickLast(username)
	}

	if err != nil && !errors.Is(err, storage.ErrNoSavedPages) {
		return e.Wrap("can not do command: can not send link", err)
	}

	if page == nil {
		return p.tg.SendMessage(chatID, msgNoSavedPages, buttonsMenu)
	}

	if err := p.tg.SendMessage(chatID, page.URL, buttonsMenu); err != nil {
		return e.Wrap("can not do command: can not send link", err)
	}

	return p.storage.Remove(page)
}

func (p *Processor) sendTag(chatID int, username string, tag string) (err error) {
	page, err := p.storage.PickTag(username, tag)
	if err != nil && !errors.Is(err, storage.ErrNoSavedPages) {
		return e.Wrap("can not do command: can not send tag", err)
	}

	if page == nil {
		return p.tg.SendMessage(chatID, msgNoSavedPages, "")
	}

	if err := p.tg.SendMessage(chatID, page.URL, ""); err != nil {
		return e.Wrap("can not do command: can not send tag", err)
	}

	return nil
}

func (p *Processor) sendID(chatID int, userName string, ID int) error {
	page, err := p.storage.PickID(userName, ID)
	if err != nil && !errors.Is(err, storage.ErrNoSavedPages) {
		return e.Wrap("can not do command: can not send tag", err)
	}

	if page == nil {
		return p.tg.SendMessage(chatID, msgNoSavedPages, "")
	}

	if err := p.tg.SendMessage(chatID, page.URL, ""); err != nil {
		return e.Wrap("can not do command: can not send page all", err)
	}

	return p.storage.Remove(page)
}

func (p *Processor) sendAll(chatID int, username string) (err error) {
	pages, err := p.storage.PickAll(username)
	if err != nil && !errors.Is(err, storage.ErrNoSavedPages) {
		return e.Wrap("can not do command: can not send tag", err)
	}

	if pages == nil {
		return p.tg.SendMessage(chatID, msgNoSavedPages, "")
	}

	response := ""
	for _, page := range pages {
		response += fmt.Sprintf("ID: %d - %s\n", page.ID, page.URL)
	}

	if err := p.tg.SendMessage(chatID, response, ""); err != nil {
		return e.Wrap("can not do command: can not send page all", err)
	}

	return nil
}

func (p *Processor) sendHelp(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelp, "")
}

func (p *Processor) sendHello(chatID int) error {
	buttonsMenu, err := createKeyBoard(ButtonGetLink, buttonAll, ButtonHelp)
	if err != nil {
		return e.Wrap("can not do buttons", err)
	}
	return p.tg.SendMessage(chatID, msgHello, buttonsMenu)
}

func (p *Processor) choiceMethod(chatID int) error {
	buttonsChoiceMethod, err := createKeyBoard(buttonLast, buttonRnd, buttonFirst)
	if err != nil {
		return e.Wrap("can not do buttons", err)
	}
	return p.tg.SendMessage(chatID, msgChoiceMethod, buttonsChoiceMethod)
}

func isAddCmd(text string) bool {
	return isURL(text)
}

func isURL(text string) bool {
	u, err := url.Parse(text)
	return err == nil && u.Host != ""
}
