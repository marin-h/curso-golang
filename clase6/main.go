package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Item struct {
	Id    int `json:"id"`
	Price int `json:"price"`
}

type Cart struct {
	Items map[int]Item
}

// Endpoints
// POST /items con payload '{"id": 1, "price": 100}'
// GET /items/1 con payload '{"id": 1, "price": 100}'
func main() {
	cart := Cart{Items: map[int]Item{}}

	http.HandleFunc("/items/", func(w http.ResponseWriter, r *http.Request) {

		item := Item{}

		if r.Method == http.MethodGet {
			id, err := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if err := json.NewEncoder(w).Encode(cart.Items[id]); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		if r.Method == http.MethodPost {
			if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			cart.Items[item.Id] = item
			w.WriteHeader(http.StatusCreated)
		}
	})
	log.Println("Escuchando en 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
