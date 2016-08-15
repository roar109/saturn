package saturn

import (
	"flag"
	"log"

	"golang.org/x/net/context"

	"google.golang.org/cloud/pubsub"
)

var (
	projID  = flag.String("p", "", "Project ID")
	client_ *pubsub.Client
)

func init() {
	//TODO validate the project id
	client, err := pubsub.NewClient(context.Background(), *projID)

	if err != nil {
		log.Println("Error creating pubsub client")
	} else {
		log.Print("GPub configured")
		client_ = client
	}
}

func checkTopicExists(client *pubsub.Client, topicName string) bool {
	exists, err := client.Topic(topicName).Exists(context.Background())
	if err != nil {
		log.Fatalf("Checking topic exists failed: %v", err)
	}
	return exists
}

func sendMessage(topicName string, message string) {
	if checkTopicExists(client_, topicName) {
		log.Print("Sending message")

		_, erro := client_.Topic(topicName).Publish(context.Background(), &pubsub.Message{
			Data: []byte(message),
		})

		if erro != nil {
			log.Fatalf("%s", erro)
		}
	} else {
		log.Fatalf("Topic does not exists")
	}
}
