package main

import (
	"fmt"
	"github.com/Armingodiz/URL-Shortener/api"
	"github.com/Armingodiz/URL-Shortener/db"
	//	"github.com/go-redis/redis"
	"log"
	"net/http"
	"time"
)

func main() {
	db := db.GetNewDatabase(8282)
	api.SetShortener(db)
	users, _ := db.RsDb.HGetAll("users").Result()

	fmt.Println(users)

	go func() {
		for {
			time.Sleep(5 * time.Second)
			sessions, _ := db.RsDb.HGetAll("sessions").Result()
			fmt.Println(sessions)
		}
	}()
	log.Fatal(http.ListenAndServe(":8080", api.GetRouter()))
}
