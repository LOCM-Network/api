package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/locm-team/api/database"
	"github.com/locm-team/api/router"
	"github.com/locm-team/api/util"
)

func main() {
	r := mux.NewRouter()
	l := log.Default()
	database.SetUp(l)
	router.SetupEndPoints(r)
	util.InitConfig()
	fmt.Println("Server is running on port " + util.GetConfig()["port"])
	log.Fatal(http.ListenAndServe(":"+util.GetConfig()["port"], r))

}
