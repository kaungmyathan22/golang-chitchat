package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kaungmyathan22/golang-chitchat/internal/handlers"
)

func main() {
	mux := routes()

	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()

	log.Println("Staring web server on port 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
	}

}
