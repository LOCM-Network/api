package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/locm-team/api/util"
)

func SetupEndPoints(r *mux.Router) {
	r.HandleFunc("/api/players", getAllPlayersHandler).Methods("GET")
	r.HandleFunc("/api/player/{name}", getPlayerHandler).Methods("GET")
	r.HandleFunc("/api/callback", callbackHandler).Methods("GET")
	r.HandleFunc("/api/register", postRegisterPlayerHandler).Methods("POST")
	r.HandleFunc("/api/donate", postCardHandler).Methods("POST")
}

func checkIP(r *http.Request) bool {
	access_ip := util.GetConfig().RemoteHost
	return util.GetIP(r) != access_ip
}
