package saturn

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//GetMessageByIDHandler handles the router for searching a message
func GetMessageByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	messageID := vars["messageId"]
	var mess = storage.read(messageID)
	fmt.Fprintf(w, mess)
}
