package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type JHandler struct {
	job Job
}

type SHandler struct {
	job Job
}

type Handler interface {
	handle(job Job)
}

func JobRouteHandler(w http.ResponseWriter, r *http.Request) {
	JobRoutesHandler(w, r, JHandler{})
}

func SJobRouteHandler(w http.ResponseWriter, r *http.Request) {
	JobRoutesHandler(w, r, SHandler{})
}

func JobRoutesHandler(w http.ResponseWriter, r *http.Request, handler Handler) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println("error reading body")
		w.WriteHeader(http.StatusBadRequest)
	}

	dec := json.NewDecoder(strings.NewReader(string(body)))
	w.WriteHeader(http.StatusOK)

	for {
		var m Job
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Println("error decoding body")
			w.WriteHeader(http.StatusBadRequest)
			break
		}
		//Generate message id
		m.buildMessageId()
		//Return the message id
		fmt.Fprintf(w, m.MessageId)

		handler.handle(m)
	}
}

func (jHandler JHandler) handle(job Job) {
	jHandler.job = job
	fmt.Println(fmt.Sprintf("Name: %s , message: %s", jHandler.job.Name, jHandler.job.Payload))
	//call job.name on remote

}
func (sHandler SHandler) handle(job Job) {
	sHandler.job = job
	fmt.Println(fmt.Sprintf("Name: %s, message: %s, Pattern: %s", sHandler.job.Name, sHandler.job.Payload, sHandler.job.Pattern))
	//schedule job on pattern
}
