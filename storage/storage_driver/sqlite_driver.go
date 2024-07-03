package storage_driver

import (
	"database/sql"
	"log"

	"github.com/anatolyjudov/job/app/model"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteDriver struct{}

func (d SqliteDriver) Init() bool {
	initDB("/Users/volganian/.job/test.db")
	return true
}

func (d SqliteDriver) SaveUser(user model.User) bool {
	return true
}

var DB *sql.DB

func initDB(filepath string) {
	var err error
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
	statement.Exec()
}
