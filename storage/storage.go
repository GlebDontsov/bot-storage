package storage

import "errors"

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	PickLast(userName string) (*Page, error)
	PickFirst(userName string) (*Page, error)
	PickTag(userName string, tag string) (*Page, error)
	PickAll(userName string) ([]*Page, error)
	PickID(userName string, id int) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

type Page struct {
	ID       int
	URL      string
	UserName string
}

var ErrNoSavedPages = errors.New("no saved pages")
