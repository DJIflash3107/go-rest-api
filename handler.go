package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handlerGetUser(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, users)
}
func handlerAddUser(w http.ResponseWriter, r *http.Request) {
	type Parameters struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
	}
	decoder := json.NewDecoder(r.Body)
	params := Parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	users = append(users, User(params))
	respondWithJSON(w, 200, struct{}{})
}
