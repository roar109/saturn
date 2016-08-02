package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/roar109/saturn"
)

func main() {
	log.Println("[Saturn]")
	handleProgramExit()

	flag.Parse()

	log.Fatal(http.ListenAndServe(":8088", saturn.NewConfiguredRouter()))
}

func handleProgramExit() {
	go func() {
		sigchan := make(chan os.Signal, 10)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		log.Println("Program killed, closing everything..")

		saturn.CloseDB()

		log.Println("Bon voyage!")
		os.Exit(0)
	}()
}
