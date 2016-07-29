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

type JobCallb func(job Job)

func JobInterface(w http.ResponseWriter, r *http.Request) {
	JInterface(w, r, func(job Job) { JobHandler(job) })
}

func SJobInterface(w http.ResponseWriter, r *http.Request) {
	JInterface(w, r, func(job Job) { SJobHandler(job) })
}

func JInterface(w http.ResponseWriter, r *http.Request, callback JobCallb) {
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

		callback(m)
	}
}

func JobHandler(job Job) {
	fmt.Println(fmt.Sprintf("Name: %s , message: %s", job.Name, job.Payload))
}

func SJobHandler(job Job) {
	fmt.Println(fmt.Sprintf("Name: %s, message: %s, Pattern: %s", job.Name, job.Payload, job.Pattern))
}
