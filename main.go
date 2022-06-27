package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/locm-team/api/database"
	"github.com/locm-team/api/router"
)

func main() {
	r := mux.NewRouter()
	l := log.Default()
	database.SetUp(l)
	router.SetupEndPoints(r)
	fmt.Println("Starting WebServer on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
