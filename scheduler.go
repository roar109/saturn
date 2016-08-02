package saturn

import (
	"encoding/json"
	"log"
)

type JobPayload struct {
	Pattern string
	Message string
}

type SJob struct {
	Key     string
	Payload JobPayload
}

func (sJob *SJob) toString() string {
	b, err := json.Marshal(sJob.Payload)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
