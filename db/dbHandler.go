package db

import (
	"github.com/go-redis/redis"
	"log"
	"strconv"
)

type DataBase struct {
	RsDb *redis.Client
}

func GetNewDatabase(port int) *DataBase {
	addr := "localhost:" + strconv.Itoa(port)
	redisClient := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatalln(err)
	}
	return &DataBase{
		RsDb: redisClient,
	}
}
func (db *DataBase) AddSession(uuidCode, username string) error {
	return nil
}
func (db *DataBase) GetSessionInfo(uuidCode string) string {
	return "user"
}
func (db *DataBase) SignUp(username, password string) error {
	// TODO check if user already exists
	// TODO check if there is a session for this user
	// TODO add user to db
	return nil
}
func (db *DataBase) Login(username, password string) error {
	return nil
}
func (db *DataBase) AddLink(originLink, shortenLink string) error {
	return nil
}
func (db *DataBase) GetLink(shortenLink string) string {
	return ""
}
func (db *DataBase) GetUrls(username string) map[string]string {
	return nil
}
func (db *DataBase) Logout(uuidCode string) error {
	// TODO remove session from database
	return nil
}
