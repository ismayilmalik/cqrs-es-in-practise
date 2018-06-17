package models

import (
	"database/sql"
)

type Article struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	AuthorID int `json:"authorId"`
}

func (a *Article) Create(db *sql.DB) error {
	stmt, err := db.Prepare("insert into articlePortal.articles(title,content,author_id) values(?,?,?) ;")
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(a.Title, a.Content, a.AuthorID)
	return err
}

func (a *Article) Edit(db *sql.DB) error {
	stmt, err := db.Prepare("update articlePortal.articles set content=?, title=? where id=? ;")
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(a.Content, a.Title, a.ID)
	return err
}

func (a *Article) Delete(db *sql.DB) error {
	stmt, err := db.Prepare("delete from articlePortal.articles set where id=? ;")
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(a.ID)
	return err
}

