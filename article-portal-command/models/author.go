package models

import (
	"database/sql"
)

type Author struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Articles []Article `json:"articles"`
}

func (a *Author) Create(db *sql.DB) error {
	stmt, err := db.Prepare("insert into articlePortal.authors(name) values(?) ;")
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(a.Name)
	return err
}

func (a *Author) Edit(db *sql.DB) error {
	stmt, err := db.Prepare("update articlePortal.authors set name=? where id=? ;")
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(a.Name, a.ID)
	return err
}

func (a *Author) Delete(db *sql.DB) error {
	stmt, err := db.Prepare("delete from articlePortal.authors set where id=? ;")
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(a.ID)
	return err
}

