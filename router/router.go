package router

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/locm-team/api/database"
	"github.com/locm-team/api/donate"
	"github.com/locm-team/api/player"
)

func SetupEndPoints(r *mux.Router) {
	r.HandleFunc("/player/{name}", getPlayerHandler).Methods("GET")
	r.HandleFunc("/register", postPlayerHandler).Methods("POST")
	r.HandleFunc("/donate", postCardHandler).Methods("POST")
}

func getPlayerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	name := vars["name"]
	if ok := database.GetDataBase().CheckPlayer(name); ok {
		playerData, ok2 := database.GetDataBase().GetPlayerData(name)
		if ok2 {
			json.NewEncoder(w).Encode(Response{Status: ResponseStatusOK, Message: ResponseOkMessage, Data: playerData})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		json.NewEncoder(w).Encode(Response{Status: ResponseStatusNotFound, Message: ResponseNotFoundMessage, Data: nil})
	}
}

func postPlayerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var playerData player.PlayerData
	err := json.NewDecoder(r.Body).Decode(&playerData)
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(Response{Status: ResponseStatusInternalServerError, Message: ResponseInternalServerErrorMessage, Data: nil})
		return
	}
	name := playerData.Name
	join_date := playerData.JoinDate
	coin := playerData.Coin
	ok2 := database.GetDataBase().Register(name, join_date, coin)
	if ok2 {
		json.NewEncoder(w).Encode(Response{Status: ResponseStatusOK, Message: ResponseOkMessage, Data: nil})
		return
	}
	json.NewEncoder(w).Encode(Response{Status: ResponseStatusInternalServerError, Message: ResponseInternalServerErrorMessage, Data: nil})
}

func postCardHandler(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query()
	telco := data.Get("telco")
	pin := data.Get("pin")
	serial := data.Get("serial")
	transaction_id := data.Get("transaction_id")
	amount, ok := strconv.ParseInt(data.Get("amount"), 10, 32)
	if ok != nil {
		amount = 0
	}
	player := data.Get("player")
	card := &donate.CardData{
		Telco:         telco,
		Pin:           pin,
		Serial:        serial,
		Amount:        amount,
		Player:        player,
		TransactionID: transaction_id,
	}
	card.PostCard()
}
