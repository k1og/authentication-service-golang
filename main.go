package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	connectionString := "mongodb://mongodb-rs-node-1:27017,mongodb-rs-node-2:27017,mongodb-rs-node-2:27017/?replicaSet=rs0"

	client, err := connectToDatabase(connectionString)
	ctx := context.TODO()

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	http.HandleFunc("/1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/1")
	})

	http.HandleFunc("/2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/2")
	})

	http.HandleFunc("/3", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/3")
	})

	http.HandleFunc("/4", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/4")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func connectToDatabase(cntnStr string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cntnStr))
	return client, err
}
