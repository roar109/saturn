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
	router.HandleFunc("/", Index)                                                                                         //.Schemes("https")
	router.HandleFunc("/message/{messageId}", GetMessageByIDHandler)                                                      //.Schemes("https")
	router.HandleFunc("/job", JobRouteHandler).Methods("POST").HeadersRegexp("Content-Type", "application/(text|json)")   //.Schemes("https")
	router.HandleFunc("/sjob", SJobRouteHandler).Methods("POST").HeadersRegexp("Content-Type", "application/(text|json)") //.Schemes("https")
	log.Fatal(http.ListenAndServe(":8088", router))
}

// Index is a simple health service
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ping pong, hello there!")
}
