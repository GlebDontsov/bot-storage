package sqlite

import (
	"bot-storage/lib/e"
	"bot-storage/storage"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(path string) (*Storage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, e.Wrap("can not open  database", err)
	}

	if err := db.Ping(); err != nil {
		return nil, e.Wrap("can not connect to database", err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Save(p *storage.Page) error {
	q := `INSERT INTO pages (url, user_name) VALUES (?, ?)`

	if _, err := s.db.Exec(q, p.URL, p.UserName); err != nil {
		return e.Wrap("can not save page", err)
	}

	return nil
}

func (s *Storage) PickRandom(userName string) (*storage.Page, error) {
	q := `SELECT url FROM pages WHERE user_name = ? ORDER BY RANDOM() LIMIT 1`

	var url string

	err := s.db.QueryRow(q, userName).Scan(&url)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, e.Wrap("can not pick random page", err)
	}

	return &storage.Page{
		URL:      url,
		UserName: userName,
	}, nil
}

func (s *Storage) Remove(page *storage.Page) error {
	q := `DELETE FROM pages WHERE url = ? AND user_name = ?`
	if _, err := s.db.Exec(q, page.URL, page.UserName); err != nil {
		return e.Wrap("can not remove page", err)
	}

	return nil
}

func (s *Storage) IsExists(page *storage.Page) (bool, error) {
	q := `SELECT COUNT(*) FROM pages WHERE url = ? AND user_name = ?`

	var count int

	if err := s.db.QueryRow(q, page.URL, page.UserName).Scan(&count); err != nil {
		return false, e.Wrap("can not check if page exists", err)
	}

	return count > 0, nil
}

func (s *Storage) Init() error {
	q := `CREATE TABLE IF NOT EXISTS pages (url TEXT, user_name TEXT)`

	_, err := s.db.Exec(q)
	if err != nil {
		return e.Wrap("can not create table", err)
	}
	return nil
}
