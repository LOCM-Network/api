package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/locm-team/api/database"
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
	r.HandleFunc("/player/register/{data}", postPlayerHandler).Methods("POST")
}

func getPlayerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	name := vars["name"]
	database := &database.PlayerData{
		Name: name,
	}
	joinDate := database.GetJoinDate()
	json.NewEncoder(w).Encode(joinDate)
}

func postPlayerHandler(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query()
	name := data.Get("name")
	join_date := data.Get("join_date")

	database := &database.PlayerData{
		Name:     name,
		JoinDate: join_date,
	}
	ok := database.Register()
	if ok {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}
