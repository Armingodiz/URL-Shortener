package main

import (
	"fmt"
	"github.com/Armingodiz/URL-Shortener/api"
	"github.com/Armingodiz/URL-Shortener/db"
	"log"
	"net/http"
	"time"
)

func main() {
	db := db.GetNewDatabase(8282)
	api.SetShortener(db)
	log.Fatal(http.ListenAndServe(":8080", api.GetRouter()))
}
