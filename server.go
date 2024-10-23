package main

import (
	"fmt"
	"log"
	"net/http"
)

//http://localhost:4242/create-payment-intent
//http://localhost:4242/health

func main() {
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)

	log.Println("Listening on localhost:4242...")
	var err error = http.ListenAndServe("localhost:4242", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENDPOINT CALLED!")

}
