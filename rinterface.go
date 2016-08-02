package saturn

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//JHandler is a representation os a simple job
type JHandler struct {
	job Job
}

//SHandler is a representation os a simple scheduled job
type SHandler struct {
	job Job
}

//Handler represents generic contract of how a job should be handled by a router
type Handler interface {
	handle(job Job)
}

//JobRouteHandler handle the router for the Job
func JobRouteHandler(w http.ResponseWriter, r *http.Request) {
	jobRoutesHandler(w, r, JHandler{})
}

//SJobRouteHandler handle the router for the scheduled Job
func SJobRouteHandler(w http.ResponseWriter, r *http.Request) {
	jobRoutesHandler(w, r, SHandler{})
}

//jobRoutesHandler is generic method to handle any valid Handler value
func jobRoutesHandler(w http.ResponseWriter, r *http.Request, handler Handler) {
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
		//TODO Error handling
		//TODO check if we do this async/sync
		handler.handle(m)
	}
}

func (jHandler JHandler) handle(job Job) {
	jHandler.job = job
	log.Println(fmt.Sprintf("Name: %s , payload: %s, msgid: %s", jHandler.job.Name, jHandler.job.Payload, jHandler.job.MessageId))
	//call job.name on remote
	go func() {
		sendMessage(jHandler.job.Name, jHandler.job.Payload)
	}()
}

func (sHandler SHandler) handle(job Job) {
	sHandler.job = job
	log.Println(fmt.Sprintf("Name: %s, message: %s, Pattern: %s, msgId: %s", sHandler.job.Name, sHandler.job.Payload, sHandler.job.Pattern, sHandler.job.MessageId))

	jobPayload := &JobPayload{Pattern: sHandler.job.Pattern, Message: sHandler.job.Payload}
	var sjob = &SJob{Key: sHandler.job.MessageId, Payload: *jobPayload}
	//save the message
	storage.save(*sjob)
}
