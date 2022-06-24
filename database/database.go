package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDatabase struct {
	Database *sql.DB
	Logger   *log.Logger
}

func (db *SQLiteDatabase) SetUp() {
	log.Println("Initializing database")
	file, err := os.Create("database.db")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	sqliData, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	db.createTable()
	db.Database = sqliData
}

func (db *SQLiteDatabase) createTable() {
	createPlayerTable := `CREATE TABLE IF NOT EXISTS player	(
		"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name"	TEXT NOT NULL,
		"join_date"	TEXT NOT NULL
	);`
	statement, err := db.Database.Prepare(createPlayerTable)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
}

func (db *SQLiteDatabase) InsertPlayer(name string, joinDate string) {
	insertPlayer := `INSERT INTO player (name, join_date) VALUES (?, ?)`
	statement, err := db.Database.Prepare(insertPlayer)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec(name, joinDate)
}

func (db *SQLiteDatabase) GetJoinDate(name string) string {
	getJoinDate := `SELECT join_date FROM player WHERE name = ?`
	statement, err := db.Database.Prepare(getJoinDate)
	if err != nil {
		log.Fatal(err)
	}
	var joinDate string
	statement.QueryRow(name).Scan(&joinDate)
	return joinDate
}
