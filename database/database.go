package database

import (
	"os"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	//"github.com/kennabila/kalamadali/entity"
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
	id := ""
	err := db.Database.QueryRow("SELECT telegram_id from users where telegram_id=?", telegram_id).Scan(&id)

	if err != nil {
		_, err = db.Database.Exec("INSERT INTO users (telegram_id, github_id, deleted) VALUES ('" + telegram_id + "', '" + github_id + "', 0)")
	} else {
		_, err = db.Database.Exec("UPDATE users SET github_id=?, deleted=? where telegram_id=?", github_id, 0, telegram_id)
	}

	//user := &entity.User{
	//	TelegramId: telegram_id,
	//	GithubId: github,
	//	Deleted: 0,
	//}

	checkErr(err)
}

func (db *Database) Delete(telegram_id string) {
	_, _ = db.Database.Exec("UPDATE users SET deleted=? where telegram_id=?", 1, telegram_id)
}

func (db *Database) Update(telegram_id string, github_id string) string {
	var id, deleted string
	err := db.Database.QueryRow("SELECT telegram_id, deleted from users where telegram_id=?", telegram_id).Scan(&id, &deleted)

	if err != nil {
		return "not_found"
	} else {
		if deleted == "1" {
			return "deleted"
		} else {
			_, err = db.Database.Exec("UPDATE users SET github_id=? where telegram_id=?", github_id, telegram_id)
			checkErr(err)
		}
	}

	return "succeed"
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
