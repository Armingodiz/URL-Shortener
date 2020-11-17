package db

import (
	"errors"
	"fmt"
)

//////////////////////////////////////////////////// adding session :

func (db *DataBase) AddSession(uuidCode, username string) error {
	err := db.RsDb.HSet("sessions", uuidCode, username).Err()
	if err != nil {
		return errors.New("ERROR IN ADDING SESSION TO database")
	}
	return nil
}

//////////////////////////////////////////////////// getting username fore session :

func (db *DataBase) GetSessionInfo(uuidCode string) string {
	userName, _ := db.RsDb.HGet("sessions", uuidCode).Result()
	return userName
}

//////////////////////////////////////////////////// signing up user :

func (db *DataBase) SignUp(username, password string) error {
	// check if user already exists
	user, _ := db.RsDb.HGet("users", username).Result()
	if user != "" {
		return errors.New("USER Alrady exists !")
	}
	// add user to database
	err := db.RsDb.HSet("users", username, password).Err()
	if err != nil {
		fmt.Println(err)
		return errors.New("EROR in adding user to database")
	}
	return nil
}

//////////////////////////////////////////////////////// login user :

func (db *DataBase) Login(username, password string) error {
	fmt.Println("logging in ...")
	pass, _ := db.RsDb.HGet("users", username).Result()
	if pass == "" {
		return errors.New("THIS USER DOES NOT EXIST !")
	} else if pass != password {
		fmt.Println(pass)
		fmt.Println(password)
		return errors.New("wrong password!")
	}
	if sessionExistence(db, username) {
		return errors.New("LOGGED IN from another device ! ")
	}
	return nil
}

//////////////////////////////////////////////////////  logout user :

func (db *DataBase) Logout(uuidCode string) error {
	err := db.RsDb.HDel("sessions", uuidCode)
	if err != nil {
		return errors.New("error in removing session from database")
	}
	return nil
}

//////////////////////////////////////////////////////// check existence of session :

func sessionExistence(db *DataBase, username string) bool {
	sessions, _ := db.RsDb.HGetAll("sessions").Result()
	for _, userN := range sessions {
		if userN == username {
			return true
		}
	}
	return false
}
