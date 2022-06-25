package database

import "log"

type PlayerData struct {
	Name     string
	JoinDate string
	Coin     int64
	Database *SQLiteDatabase
	Log      *log.Logger
}

type PlayerReturn struct {
	Name     string `json:"name"`
	JoinDate string `json:"join_date"`
	Coin     int64  `json:"coin"`
}

func (p *PlayerData) GetJoinDate() (join string, ok bool) {
	if join, ok = p.Database.GetJoinDate(p.Name); ok {
		return join, true
	}
	return "", false
}

func (p *PlayerData) GetCoin() (coin int64, ok bool) {
	if coin, ok = p.Database.GetCoin(p.Name); !ok {
		p.Log.Println("Error getting coin")
		return 0, false
	}
	return coin, true
}

func (p *PlayerData) Register() bool {
	p.Database.InsertPlayer(p.Name, p.JoinDate, p.Coin)
	p.Log.Fatal("Register player: " + p.Name)
	return true
}

func (p *PlayerData) AddCoin(coin int64) bool {
	current, ok := p.GetCoin()
	if !ok {
		return false
	}
	p.Database.SetCoin(p.Name, current+coin)
	return true
}

func (p *PlayerData) RemoveCoin(coin int64) bool {
	current, ok := p.GetCoin()
	if !ok {
		return false
	}
	p.Database.SetCoin(p.Name, current-coin)
	return true
}

//get PlayerReturn
func (p *PlayerData) GetPlayerReturn() PlayerReturn {
	return PlayerReturn{
		Name:     p.Name,
		JoinDate: p.JoinDate,
		Coin:     p.Coin,
	}
}
