package db

import (
	//"errors"
	"github.com/go-redis/redis"
	"strconv"
	"log"
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



////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (db *DataBase) AddLink(originLink, shortenLink,userName string) error {
	return nil
}
func (db *DataBase) GetLink(shortenLink string) string {
	return ""
}
func (db *DataBase) GetUrls(username string) map[string]string {
	return nil
}
