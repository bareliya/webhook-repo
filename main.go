package main

import (
	"github.com/webhook-repo/server"
	"github.com/webhook-repo/ui"
	"os"
	"os/signal"
	"log"
	"syscall"
)

func main() {
    go server.StartServer()
	ui.StartUI()


	// Listen for signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	
	// Select on error channels from different modules
	for {
		select {
			case sig := <-sigs:
				log.Println("Got signal, beginning shutdown %s", sig)
				os.Exit(1)
		}
	}

}
