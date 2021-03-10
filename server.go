package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Mongo
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongo-service:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := client.Ping(context.TODO(), nil)
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte(fmt.Sprintf("Hello World")))
	})

	server := http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	fmt.Println("Server is listening on port :8000")
	server.ListenAndServe()
}
