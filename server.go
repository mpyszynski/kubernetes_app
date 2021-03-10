package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	testDB := client.Database("test")
	testColl := testDB.Collection("test-col")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := client.Ping(context.TODO(), nil)
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte(fmt.Sprintf("Hello World")))
	})

	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		cursor, err := testColl.Find(context.TODO(), bson.M{})
		if err != nil {
			log.Fatal(err)
		}
		var results []bson.M
		if err = cursor.All(context.TODO(), &results); err != nil {
			log.Fatal(err)
		}
		w.Write([]byte(fmt.Sprintf("Inserted docs %v", results)))
	})

	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		res, err := testColl.InsertOne(context.TODO(), bson.D{
			{Key: "name", Value: "Dyzio"},
		})
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte(fmt.Sprintf("Inserted docs %v", res.InsertedID)))
	})

	server := http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	fmt.Println("Server is listening on port :8000")
	server.ListenAndServe()
}
