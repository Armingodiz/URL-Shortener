package main

import (
	"github.com/Armingodiz/URL-Shortener/api"
	"github.com/Armingodiz/URL-Shortener/db"
	"log"
	"net/http"
	"fmt"
)

func main() {
	db := db.GetNewDatabase(8282)
	api.SetShortener(db)
	db.RsDb.HSet("last", "armin", "godarzi")
	db.RsDb.HSet("last", "hadi", "abbasi")
	fmt.Println(db.RsDb.HGetAll("last"))
	log.Fatal(http.ListenAndServe(":8080", api.GetRouter()))
}
