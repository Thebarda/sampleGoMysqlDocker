package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"./paths"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/fetch/user", path.User).Methods("GET")
	router.HandleFunc("/insert/user", path.InsertUser).Methods("POST")
	fmt.Println("HTTP server listening on port " + os.Getenv("GO_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("GO_PORT"), router))
}
