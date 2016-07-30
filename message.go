package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//GetMessageByIDHandler handles the router for searching a message
func GetMessageByIDHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Checking ")
	vars := mux.Vars(r)
	messageID := vars["messageId"]
	fmt.Fprintf(w, messageID)
}
