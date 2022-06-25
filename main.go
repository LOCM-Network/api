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
)

func main() {
	r := mux.NewRouter()
	l := log.Default()
	sqlite := database.SQLiteDatabase{
		Logger: l,
	}
	sqlite.SetUp()
	setupEnpoints(r)
	fmt.Println("Starting WebServer on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func setupEnpoints(r *mux.Router) {
	r.HandleFunc("/player/{name}", getPlayerHandler).Methods("GET")
	r.HandleFunc("/player/register", postPlayerHandler).Methods("POST")
	r.HandleFunc("/donate", postCardHandler).Methods("POST")
}

func getPlayerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	name := vars["name"]
	database := &database.PlayerData{
		Name: name,
	}
	playerdata := database.GetPlayerReturn()
	e, err := json.Marshal(playerdata)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(e)
}

func postPlayerHandler(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query()
	name := data.Get("name")
	join_date := data.Get("join_date")
	coin, ok := strconv.ParseInt(data.Get("coin"), 10, 32)
	if ok != nil {
		coin = 0
	}
	database := &database.PlayerData{
		Name:     name,
		JoinDate: join_date,
		Coin:     coin,
	}
	ok2 := database.Register()
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
