package handlers

import (
	"database/sql"
	"net/http"
)

type ArticleHandler struct {
	DB *sql.DB
}

func (a *ArticleHandler) Add(w http.ResponseWriter, r *http.Request){

}

func (a *ArticleHandler) Update(w http.ResponseWriter, r *http.Request){

}

func (a *ArticleHandler) Remove(w http.ResponseWriter, r *http.Request){

}