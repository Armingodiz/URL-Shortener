package db

import (
	"errors"
	"github.com/go-redis/redis"
	"log"
	"strconv"
)

type DataBase struct {
	RsDb *redis.Client
}

///////////////////////////////////////////////////// new redis database :

func GetNewDatabase(port int) *DataBase {
	addr := "localhost:" + strconv.Itoa(port)
	redisClient := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatalln(err)
	}
	redisClient.Del("sessions")
	return &DataBase{
		RsDb: redisClient,
	}
}

/////////////////////////////////////////////////////////// add shorten link and its original version to database

func (db *DataBase) AddLink(originLink, shortenLink, userName string) error {
	err := db.RsDb.HSet(userName, shortenLink, originLink).Err()
	if err != nil {
		return errors.New("ERROR IN ADDING SHORTEN LINK TO user  database")
	}
	err = db.RsDb.HSet("urls", shortenLink, originLink).Err()
	if err != nil {
		return errors.New("ERROR IN ADDING SHORTEN LINK TO database")
	}
	return nil
}

///////////////////////////////////////////////////////////////// get original link from database :

func (db *DataBase) GetLink(shortenLink string) string {
	url, _ := db.RsDb.HGet("urls", shortenLink).Result()
	return url
}

////////////////////////////////////////////////////////////////////// get all urls that each user have :

func (db *DataBase) GetUrls(username string) map[string]string {
	urls, _ := db.RsDb.HGetAll(username).Result()
	return urls
}
