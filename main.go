package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/locm-team/api/database"
	"github.com/locm-team/api/donate"
	"github.com/locm-team/api/player"
)

func main() {
	r := mux.NewRouter()
	l := log.Default()
	database.SetUp(l)
	setupEnpoints(r)
	fmt.Println("Starting WebServer on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func setupEnpoints(r *mux.Router) {
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
			json.NewEncoder(w).Encode(playerData)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Player not found"})
	}
}

func postPlayerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var playerData player.PlayerData
	err := json.NewDecoder(r.Body).Decode(&playerData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	name := playerData.Name
	join_date := playerData.JoinDate
	coin := playerData.Coin
	ok2 := database.GetDataBase().Register(name, join_date, coin)
	if ok2 {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
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
