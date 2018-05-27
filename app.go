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
	router.HandleFunc("/", path.Index)
	fmt.Println("HTTP server listening on port " + os.Getenv("GO_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("GO_PORT"), router))
}
