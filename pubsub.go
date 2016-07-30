package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"golang.org/x/net/context"

	"google.golang.org/cloud/pubsub"
)

var (
	projID     = flag.String("p", "", "The ID of your Google Cloud project.")
	subName    = flag.String("s", "", "The name of the subscription to pull from")
	numConsume = flag.Int("n", 10, "The number of messages to consume")
)

func main1() {
	flag.Parse()

	if *projID == "" {
		log.Fatal("-p is required")
	}
	if *subName == "" {
		log.Fatal("-s is required")
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	fmt.Println("pass signal")
	ctx := context.Background()
	fmt.Println("pass background")
	client, err := pubsub.NewClient(ctx, *projID)
	if err != nil {
		log.Fatalf("creating pubsub client: %v", err)
	}
	fmt.Println("pass new client")
	sub := client.Subscription(*subName)
	fmt.Println("pass subscription")
	it, err := sub.Pull(ctx, pubsub.MaxExtension(time.Minute))
	if err != nil {
		fmt.Printf("error constructing iterator: %v", err)
		return
	}
	defer it.Stop()

	go func() {
		<-quit
		it.Stop()
	}()
	fmt.Println("before loop")
	for i := 0; i < *numConsume; i++ {
		m, err := it.Next()
		if err == pubsub.Done {
			break
		}
		if err != nil {
			fmt.Printf("advancing iterator: %v", err)
			break
		}
		fmt.Printf("got message: %v\n", string(m.Data))
		m.Done(true)
	}
}
