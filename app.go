package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Account struct {
	ID string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	UserName string `json:"user_name"`
}

var accountMap map[string]Account

func init() {
	accountMap = make(map[string]Account)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Ping")
	})
	r.HandleFunc("/accounts", createAccountHandler).Methods("POST")
	r.HandleFunc("/accounts/{id}", getAccountHandler).Methods("GET")
	r.HandleFunc("/accounts/{id}", deleteAccountHandler).Methods("DELETE")

	log.Print("Server listening on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func createAccountHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Request received to create an Account")

	var account Account
	json.NewDecoder(r.Body).Decode(&account)
	id := account.ID
	accountMap[id] = account

	log.Print("Added the Account ", account, " to list of accounts ", accountMap)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

func getAccountHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	log.Print("Request received to get and Account by account id: ", id)

	account, ok := accountMap[id]
	w.Header().Add("Content-Type", "application/json")

	if ok {
		log.Print("Successfully retrieved the account ", account, " for account id: ", id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(account)
	} else {
		log.Print("Requested account is not found for account id: ", id)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w)
	}
}

func deleteAccountHandler(w http.ResponseWriter, r * http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	log.Print("Request received to delete an Account by account id: ", id)

	_, present := accountMap[id]

	if present {
		log.Print("Removing account ", id)
		delete(accountMap, id)
		w.WriteHeader(http.StatusOK)
	} else {
		log.Print("Request to delete account failed: account ", id, " does not exist")
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w)
}

