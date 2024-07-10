package storage_driver

import (
	"database/sql"
	"log"

	"github.com/anatolyjudov/job/app/model"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteDriver struct{}

func (d SqliteDriver) GetUserById(id int) *model.User {
	var name, avatar string

	initDB("/Users/volganian/.job/test.db")

	err := DB.QueryRow("select name, avatar from user where id = ?", id).Scan(&name, &avatar)

	if err != nil {
		log.Fatalf("Failed to select user: %v", err)
	}

	u := model.User{}
	u.SetId(id).SetName(name).SetAvatarFromString(avatar)

	return &u
}

func (d SqliteDriver) GetUsers() []*model.User {

	initDB("/Users/volganian/.job/test.db")

	rows, err := DB.Query("select id, name, avatar from user")

	if err != nil {
		log.Fatalf("Failed to select usesr: %v", err)
	}

	defer rows.Close()

	var users []*model.User

	for rows.Next() {
		var id int
		var name, avatar string

		err := rows.Scan(&id, &name, &avatar)
		if err != nil {
			log.Fatal(err)
		}

		u := model.User{}
		u.SetId(id).SetName(name).SetAvatarFromString(avatar)

		users = append(users, &u)
	}

	return users
}

func (d SqliteDriver) DeleteUser(user *model.User) bool {
	initDB("/Users/volganian/.job/test.db")

	if user.Id() == 0 {
		return true
	}

	_, err := DB.Exec("delete from user where id = ?", user.Id())

	if err != nil {
		log.Fatalf("Failed to delete user: %v", err)
	}

	return true
}

func (d SqliteDriver) SaveUser(user *model.User) bool {

	initDB("/Users/volganian/.job/test.db")

	if user.Id() == 0 {
		id := d.createUser(user.Name(), user.AvatarAsString())
		user.SetId(id)
	} else {
		d.updateUser(user.Id(), user.Name(), user.AvatarAsString())
	}

	return true
}

func (d SqliteDriver) createUser(name string, avatar string) int {
	var userSql string
	var err error

	userSql = "insert into user (name, avatar) values (?, ?)"
	result, err := DB.Exec(userSql, name, avatar)

	if err != nil {
		log.Fatalf("Failed to insert user: %v", err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return int(insertId)
}

func (d SqliteDriver) updateUser(id int, name string, avatar string) {
	var userSql string
	var err error

	userSql = "update user set name = ?, avatar = ? where id = ?"
	_, err = DB.Exec(userSql, name, avatar, id)

	if err != nil {
		log.Fatalf("Failed to update user: %v", err)
	}
}

var DB *sql.DB

func initDB(filepath string) {
	var err error
	log.Output(0, "Initialising DB")
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
