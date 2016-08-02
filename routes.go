package saturn

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewConfiguredRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)                                                                                         //.Schemes("https")
	router.HandleFunc("/message/{messageId}", GetMessageByIDHandler)                                                      //.Schemes("https")
	router.HandleFunc("/job", JobRouteHandler).Methods("POST").HeadersRegexp("Content-Type", "application/(text|json)")   //.Schemes("https")
	router.HandleFunc("/sjob", SJobRouteHandler).Methods("POST").HeadersRegexp("Content-Type", "application/(text|json)") //.Schemes("https")
	return router
}

// Index is a simple health service
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ping pong, hello there!")
}
