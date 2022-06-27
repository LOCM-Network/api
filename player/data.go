package player

type PlayerData struct {
	Name     string `json:"name"`
	JoinDate string `json:"join_date"`
	Coin     int    `json:"coin"`
}

func (p *PlayerData) GetName() string {
	return p.Name
}

func (p *PlayerData) GetJoinDate() string {
	return p.JoinDate
}

func (p *PlayerData) GetCoin() int {
	return p.Coin
}

func (p *PlayerData) AddCoin(coin int) {
	p.Coin += coin
}

func (p *PlayerData) ReduceCoin(coin int) {
	p.Coin -= coin
}
