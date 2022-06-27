package router

type CardData struct {
	Telecom       string `json:"telecom"`
	Pin           string `json:"pin"`
	Serial        string `json:"serial"`
	Amount        int    `json:"amount"`
	TransactionID string `json:"transaction_id"`
	Player        string `json:"player"`
	DateTime      string `json:"date_time"`
}
