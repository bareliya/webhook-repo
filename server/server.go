package server

import (
	"fmt"
	"github.com/webhook-repo/handlers"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/webhook", handlers.WebhookHandler)
	fmt.Println("Start...")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}

}
