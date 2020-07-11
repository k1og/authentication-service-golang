package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

type AccessTokenClaims struct {
	jwt.StandardClaims
	GUID string
}

func main() {

	connectionString := "mongodb://mongodb-rs-node-1:27017,mongodb-rs-node-2:27017,mongodb-rs-node-2:27017/?replicaSet=rs0"

	client, err := connectToDatabase(connectionString)
	ctx := context.TODO()

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	secret := []byte("some secret string")

	http.HandleFunc("/1", func(w http.ResponseWriter, r *http.Request) {
		guidParam := r.URL.Query()["guid"]

		if len(guidParam) != 1 {
			panic("invalid guid")
		}

		guid := guidParam[0]

		rfTkn, err := uuid.New()

		if err != nil {
			panic(err)
		}

		RefreshToken := rfTkn[:]

		AccessToken, err := generateAccessToken(AccessTokenClaims{
			GUID: guid,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
			},
		}).SignedString(secret)

		if err != nil {
			panic(err)
		}

		fmt.Println("/1")
		fmt.Println("AccessToken", AccessToken)
		fmt.Println("RefreshToken", base64.StdEncoding.EncodeToString(RefreshToken))
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

func generateRefreshToken() (uuid.UUID, error) {
	return uuid.New()
}

func generateAccessToken(claims AccessTokenClaims) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
}
