package main

import (
	"github.com/Armingodiz/URL-Shortener/api"
	"github.com/Armingodiz/URL-Shortener/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	r := mux.NewRouter()

	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)

	// Bind to a port and pass our router in
	db := db.GetNewDatabase(8282)
	api.SetShortener(db)
	//fmt.Println(db.RsDb.Get("armin"))
	log.Fatal(http.ListenAndServe(":8000", api.GetRouter()))
}
