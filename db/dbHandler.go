package db
import "fmt"
type Database struct {
	ss string
}

func (db *Database) PrintString() {
	fmt.Println(db.ss)
}

func GetNewDatabase() *Database {
	return &Database{
		ss: "fuck you",
	}
}
