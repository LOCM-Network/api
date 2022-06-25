package player

type PlayerData struct {
	Name     string `json:"name"`
	JoinDate string `json:"join_date"`
	Coin     int64  `json:"coin"`
}

func (p *PlayerData) GetName() string {
	return p.Name
}

func (p *PlayerData) GetJoinDate() string {
	return p.JoinDate
}

func (p *PlayerData) GetCoin() int64 {
	return p.Coin
}

func (p *PlayerData) AddCoin(coin int64) {
	p.Coin += coin
}

func (p *PlayerData) ReduceCoin(coin int64) {
	p.Coin -= coin
}
