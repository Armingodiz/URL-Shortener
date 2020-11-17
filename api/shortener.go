package api

import (
	"github.com/gorilla/mux"
)

var shortener *Shortener

func SetShortener(db DataBase) {
	shortener = NewShortener(db)
}

type Shortener struct {
	router *mux.Router
	db     DataBase
}

func NewShortener(db DataBase) *Shortener {
	return &Shortener{
		router: newRouter(),
		db:     db,
	}
}
func GetRouter() *mux.Router {
	return shortener.router
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", mainPage)
	r.HandleFunc("/singUp", singUp)
	r.HandleFunc("/login", login)
	r.HandleFunc("/logout", logout)
	r.HandleFunc("/new", addLink)
	r.HandleFunc("/show", showLinks)
	r.HandleFunc("/:hash", redirect)
	return r
}
