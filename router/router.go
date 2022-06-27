package router

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/locm-team/api/database"
	"github.com/locm-team/api/player"
	"github.com/locm-team/api/util"
)

func SetupEndPoints(r *mux.Router) {
	r.HandleFunc("/api/players", getAllPlayersHandler).Methods("GET")
	r.HandleFunc("/api/player/{name}", getPlayerHandler).Methods("GET")
	r.HandleFunc("/api/register", postRegisterPlayerHandler).Methods("POST")
	r.HandleFunc("/api/donate", postCardHandler).Methods("POST")
}

func getAllPlayersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(database.GetDataBase().GetAllPlayerData())
}

func getPlayerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	name := vars["name"]
	if playerData, ok := database.GetDataBase().GetPlayerData(name); ok {
		json.NewEncoder(w).Encode(Response{Status: ResponseStatusOK, Message: ResponseOkMessage, Data: playerData})
	} else {
		json.NewEncoder(w).Encode(Response{Status: ResponseStatusNotFound, Message: ResponseNotFoundMessage, Data: nil})
	}
}

func postRegisterPlayerHandler(w http.ResponseWriter, r *http.Request) {
	var playerData player.PlayerData
	err := json.NewDecoder(r.Body).Decode(&playerData)
	log.Print(playerData)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
		json.NewEncoder(w).Encode(Response{Status: ResponseStatusInternalServerError, Message: ResponseInternalServerErrorMessage, Data: nil})
	}

	name := playerData.Name
	join_date := playerData.JoinDate
	coin := playerData.Coin
	ok2 := database.GetDataBase().Register(name, join_date, coin)
	if ok2 {
		log.Printf("Register new player=" + name + " Joined=" + join_date + " coin=" + strconv.Itoa(coin))
		json.NewEncoder(w).Encode(Response{Status: ResponseStatusOK, Message: ResponseOkMessage, Data: nil})
		return
	}
	json.NewEncoder(w).Encode(Response{Status: ResponseStatusInternalServerError, Message: ResponseInternalServerErrorMessage, Data: nil})
}

func postCardHandler(w http.ResponseWriter, r *http.Request) {
	//todo: parse card data
}

func checkIP(r *http.Request) bool {
	access_ip := util.GetConfig().RemoteHost
	return util.GetIP(r) != access_ip
}
