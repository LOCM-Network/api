package database

import "log"

type PlayerData struct {
	Name     string
	JoinDate string
	Database *SQLiteDatabase
	Log      *log.Logger
}

type JoinDate struct {
	JoinDate string `json:"join_date"`
}

func (p *PlayerData) Save() {
	p.Database.InsertPlayer(p.Name, p.JoinDate)
	p.Log.Fatal("Saved player: " + p.Name)
}

func (p *PlayerData) GetJoinDate() *JoinDate {
	return &JoinDate{JoinDate: p.Database.GetJoinDate(p.Name)}
}

//insert player into database
func (p *PlayerData) Register() bool {
	p.Database.InsertPlayer(p.Name, p.JoinDate)
	p.Log.Fatal("Register player: " + p.Name)
	return true
}
