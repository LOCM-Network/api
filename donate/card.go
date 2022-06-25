package donate

type CardData struct {
	Telco         string `json:"telco"`
	Pin           string `json:"pin"`
	Serial        string `json:"serial"`
	Amount        int64  `json:"amount"`
	Player        string `json:"player"`
	TransactionID string `json:"transaction_id"`
}

func (card *CardData) PostCard() {
	// not yet implemented
}
