package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Saturn is starting ...")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/job", JobInterface).Methods("POST")
	router.HandleFunc("/sjob", SJobInterface).Methods("POST")
	log.Fatal(http.ListenAndServe(":8088", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ping pong, hello there!")
}
