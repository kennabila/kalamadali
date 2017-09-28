package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"

	"fmt"
)

type Database struct {
	Database *sql.DB
}

func NewDatabase() *Database {
	db, err := sql.Open("mysql", os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp(127.0.0.1:3306)/")
	checkErr(err)
	initDatabase(db)

	return &Database{db}
}

func (db *Database) Insert(telegram_id string, github_id string) {
	_, err := db.Database.Exec("INSERT INTO users (telegram_id, github_id, deleted) VALUES ('"+ telegram_id + "', '"+ github_id +"', 0)")
	fmt.Println()
	fmt.Println()
	fmt.Println(err)

	checkErr(err)
}

func (db *Database) Delete(telegram_id string, github_id string) {

}

func initDatabase(db *sql.DB) {
	_, _ = db.Exec("CREATE DATABASE " + os.Getenv("DB_NAME"))
	_, _ = db.Exec("USE " + os.Getenv("DB_NAME"))
	_, _ = db.Exec("CREATE TABLE users ( id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, telegram_id varchar(50), github_id varchar(50), deleted boolean )")
	_, _ = db.Exec("ALTER TABLE users ADD INDEX (telegram_id)")
	_, _ = db.Exec("ALTER TABLE users ADD INDEX (telegram_id, github_id)")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
