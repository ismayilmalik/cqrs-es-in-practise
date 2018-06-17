package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ismayilmalik/cqrs-es-in-practise/article-portal-command/handlers"
)

const (
	AUTHOR  = "author"
	ARTICLE = "article"
)

type App struct {
	Handlers map[string]handlers.AppHandler
	Router   *mux.Router
}

func (a *App) Prepare(db *sql.DB) {
	a.initHandlers(db)
	a.initRoutes()
}

func (a *App) Run(addr string) {
	s := &http.Server{
		Addr:         addr,
		Handler:      a.Router,
		WriteTimeout: 10 * time.Second,
	}
	s.ListenAndServe()
}

func (a *App) initHandlers(db *sql.DB) {
	a.Handlers = make(map[string]handlers.AppHandler, 0)
	a.Handlers[AUTHOR] = &handlers.AuthorHandler{DB: db}
	a.Handlers[ARTICLE] = &handlers.ArticleHandler{DB: db}
}

func (a *App) initRoutes() {
	a.Router = mux.NewRouter()

	a.Router.HandleFunc("/authors", a.Handlers[AUTHOR].Add).Methods("POST")
	a.Router.HandleFunc("/authors/{id:[0-9]}", a.Handlers[AUTHOR].Update).Methods("PUT")
	a.Router.HandleFunc("/authors/{id:[0-9]}", a.Handlers[AUTHOR].Remove).Methods("DELETE")

	a.Router.HandleFunc("/articles", a.Handlers[ARTICLE].Add).Methods("POST")
	a.Router.HandleFunc("/articles", a.Handlers[ARTICLE].Update).Methods("PUT")
	a.Router.HandleFunc("/articles", a.Handlers[ARTICLE].Remove).Methods("DELETE")
}
