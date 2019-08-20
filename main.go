package main

import (
	"./api"
	"./db"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	CassandraSession := db.Session
	defer CassandraSession.Close()
	router := mux.NewRouter()
	
	router.HandleFunc("/getTransactions/{address}", api.GetTransactionDetails).Methods("GET")
	router.HandleFunc("/addTransactions/", api.InsertTransactionDetails)
	http.ListenAndServe(":8000", router)
}

