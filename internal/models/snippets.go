package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `
		INSERT INTO snippets (title, content, created, expires)
		VALUES ($1, $2, NOW(), NOW() + ($3 || ' days')::interval)
		RETURNING id
	`

	var id int
	err := m.DB.QueryRow(stmt, title, content, expires).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *SnippetModel) Get(id int) (Snippet, error) {
	stmt := `SELECT title, content, created, expires FROM snippets WHERE 
	expires > NOW() AND id = $1`
	err := m.DB.QueryRow(stmt, id)

	var s Snippet
	return Snippet{}, nil
}

func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
