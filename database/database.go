package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type Database struct {
	Database *sql.DB
}

func NewDatabase() *MySql {
	db, err := sql.Open("mysql", os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp(127.0.0.1:3306)/")
	checkErr(err)

	return &MySql{db}
}

func InitDatabase(db *sql.DB) {
	_, err = db.Exec("CREATE DATABASE " + os.Getenv("DB_NAME"))
	checkErr(err)

	_, err = db.Exec("USE " + os.Getenv("DB_NAME"))
	checkErr(err)

	_, err = db.Exec("CREATE TABLE users ( id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, telegram varchar(50), github varchar(50), deleted boolean )")
	checkErr(err)
}

func checkErr(err *error) {
	if err != nil {
		panic(err)
	}
}
