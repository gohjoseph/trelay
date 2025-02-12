package main

import (
	"os"
	"os/signal"

	"github.com/btvoidx/trelay"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Custom logger
	// log := log.New().WithField("Custom Field", true)

	server := trelay.NewServer(trelay.Options{
		Addr:       "0.0.0.0:7778",
		RemoteAddr: "localhost:7777",
	})
	// SetLogger(log) // Use custom logger

	err := server.Start()
	if err != nil {
		log.Fatalf("An error occured when starting the server: %s", err.Error())
	}

	defer server.Stop()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
}
