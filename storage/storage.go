package storage

import "errors"

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	PickLast(userName string) (*Page, error)
	PickFirst(userName string) (*Page, error)
	PickTag(userName string, tag string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

type Page struct {
	URL      string
	UserName string
}

var ErrNoSavedPages = errors.New("no saved pages")
