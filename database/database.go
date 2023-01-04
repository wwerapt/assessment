package database

import (
	"database/sql"
	"log"
	"os"
)

type DB interface {
	QueryRow(query string, args ...any) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
}

type Db struct {
	DB
	IsTestMode bool
}

func InitDb() {
	db := getDatabase()
	db.init()
}

func getDatabase() *Db {
	conn, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("can't connect to database", err)
	}

	return &Db{DB: conn}
}

func (db *Db) init() {
	createTableSql := `
	CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);
	`

	if _, err := db.Exec(createTableSql); err != nil {
		log.Fatal("can't create table ", err)
	}
}
