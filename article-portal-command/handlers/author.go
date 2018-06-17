package handlers

import (
	"github.com/ismayilmalik/cqrs-es-in-practise/apputils"
	"encoding/json"
	"github.com/ismayilmalik/cqrs-es-in-practise/article-portal-command/models"
	"database/sql"
	"net/http"
)

type AuthorHandler struct {
	DB *sql.DB
}

func (a *AuthorHandler) Add(w http.ResponseWriter, r *http.Request){
	var author models.Author

	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		apputils.HTTPRespondWithError(w, http.StatusInternalServerError, err.Error())
	}

	err = author.Create(a.DB)
	if err != nil {
		apputils.HTTPRespondWithError(w, http.StatusInternalServerError, err.Error())
	} else {
		apputils.HTTPRespondWithJSON(w, http.StatusCreated, nil)
	}
}

func (a *AuthorHandler) Update(w http.ResponseWriter, r *http.Request){
	
}

func (a *AuthorHandler) Remove(w http.ResponseWriter, r *http.Request){

}