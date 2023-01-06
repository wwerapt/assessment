package expense

import (
	"database/sql"
	"log"
	"os"
)

type handler struct {
	DB *sql.DB
}

func NewHandler(db *sql.DB) *handler {
	return &handler{db}
}

func (h *handler) InitDb() {
	h.init()
}

func GetDatabase() *handler {
	conn, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("can't connect to database", err)
	}

	return NewHandler(conn)
}

func (h *handler) init() {
	createTableSql := `
	CREATE TABLE IF NOT EXISTS expenses (
		id SERIAL PRIMARY KEY,
		title TEXT,
		amount FLOAT,
		note TEXT,
		tags TEXT[]
	);
	`

	if _, err := h.DB.Exec(createTableSql); err != nil {
		log.Fatal("can't create table ", err)
	}
}
