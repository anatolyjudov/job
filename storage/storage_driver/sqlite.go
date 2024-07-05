package storage_driver

import (
	"database/sql"
	"log"

	"github.com/anatolyjudov/job/app/model"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteDriver struct{}

func (d SqliteDriver) SaveUser(user model.User) bool {
	var userSql string
	var err error

	initDB("/Users/volganian/.job/test.db")

	if user.Id() == 0 {
		userSql = "insert into user (name, avatar) values (?, ?)"
		log.Output(0, user.Name())
		log.Output(0, user.AvatarAsString())
		_, err = DB.Exec(userSql, user.Name(), user.AvatarAsString())
	} else {
		userSql = "update user set name = ?, avatar = ? where id = ?"
		_, err = DB.Exec(userSql, user.Name(), user.AvatarAsString(), user.Id())
	}

	if err != nil {
		log.Fatalf("Failed to save user: %v", err)
	}

	return true
}

var DB *sql.DB

func initDB(filepath string) {
	var err error
	log.Output(0, "initialising DB")
	DB, err = sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatalf("Failed to open the database: %v", err)
	}

	createTables()
}

func createTables() {
	createUserTableSQL := `CREATE TABLE IF NOT EXISTS user (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT,
		"avatar" TEXT
	);`
	statement, err := DB.Prepare(createUserTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	log.Output(0, "table created")
	statement.Exec()
}
