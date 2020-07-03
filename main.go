package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
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
