package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	config := util.GetConfig()
	fmt.Println("Server is running on port " + strconv.Itoa(config.Port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Port), r))
}
