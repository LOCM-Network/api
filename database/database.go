package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/locm-team/api/player"
	_ "github.com/mattn/go-sqlite3"
)

var database *SQLiteDatabase

type SQLiteDatabase struct {
	Database *sql.DB
	Logger   *log.Logger
}

func GetDataBase() *SQLiteDatabase {
	return database
}

func SetUp(log *log.Logger) *SQLiteDatabase {
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
	database = &SQLiteDatabase{sqliData, log}
	database.createTable()
	return database
}

func (db *SQLiteDatabase) GetPlayerData(name string) (player.PlayerData, bool) {
	getPlayerData := `SELECT * FROM player WHERE name = ?`
	statement, err := db.Database.Prepare(getPlayerData)
	if err != nil {
		db.Logger.Println(err)
		return player.PlayerData{}, false
	}
	var playerData player.PlayerData
	statement.QueryRow(name).Scan(&playerData.Name, &playerData.JoinDate, &playerData.Coin)
	return playerData, true
}

func (db *SQLiteDatabase) CheckPlayer(name string) bool {
	getPlayer := `SELECT * FROM player WHERE name = ?`
	statement, err := db.Database.Prepare(getPlayer)
	if err != nil {
		db.Logger.Println(err)
		return false
	}
	var playerData player.PlayerData
	statement.QueryRow(name).Scan(&playerData.Name, &playerData.JoinDate, &playerData.Coin)
	return playerData.Name != ""
}

func (db *SQLiteDatabase) createTable() {
	createPlayerTable := `CREATE TABLE IF NOT EXISTS player	(
		"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name"	TEXT NOT NULL,
		"join_date"	TEXT NOT NULL,
		"coin"	INTEGER NOT NULL
	);`
	statement, err := db.Database.Prepare(createPlayerTable)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
}

func (db *SQLiteDatabase) Register(name string, joinDate string, coin int64) bool {
	insertPlayer := `INSERT INTO player (name, join_date, coin) VALUES (?, ?, ?)`
	statement, err := db.Database.Prepare(insertPlayer)
	if err != nil {
		log.Fatal(err)
		return false
	}
	statement.Exec(name, joinDate, coin)
	return true
}

func (db *SQLiteDatabase) GetJoinDate(name string) (JoinDate string, ok bool) {
	getJoinDate := `SELECT join_date FROM player WHERE name = ?`
	statement, err := db.Database.Prepare(getJoinDate)
	if err != nil {
		log.Fatal(err)
		return "nil", false
	}
	var joinDate string
	statement.QueryRow(name).Scan(&joinDate)
	return joinDate, true
}

func (db *SQLiteDatabase) GetCoin(name string) (coin int64, ok bool) {
	getCoin := `SELECT coin FROM player WHERE name = ?`
	statement, err := db.Database.Prepare(getCoin)
	if err != nil {
		log.Fatal(err)
		return 0, false
	}
	statement.QueryRow(name).Scan(&coin)
	return coin, true
}

func (db *SQLiteDatabase) SetCoin(name string, coin int64) bool {
	setCoin := `UPDATE player SET coin = ? WHERE name = ?`
	statement, err := db.Database.Prepare(setCoin)
	if err != nil {
		log.Fatal(err)
		return false
	}
	statement.Exec(coin, name)
	return true
}

func (db *SQLiteDatabase) GetTop5() []string {
	getTop3 := `SELECT name FROM player ORDER BY coin DESC LIMIT 5`
	statement, err := db.Database.Prepare(getTop3)
	if err != nil {
		log.Fatal(err)
	}
	var name string
	var names []string
	rows, err := statement.Query()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		rows.Scan(&name)
		names = append(names, name)
	}
	return names
}

func (db *SQLiteDatabase) Save(playerData *player.PlayerData) bool {
	insertPlayer := `INSERT OR REPLACE INTO player (name, join_date, coin) VALUES (?, ?, ?)`
	statement, err := db.Database.Prepare(insertPlayer)
	if err != nil {
		log.Fatal(err)
		return false
	}
	statement.Exec(playerData.Name, playerData.JoinDate, playerData.Coin)
	return true
}
