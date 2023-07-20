package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Banker interface {
	PrintUser(w http.ResponseWriter, r *http.Request)
}

// constructor - creat user

func CreateBank() *Bank {
	return &Bank{
		Id:   rand.Intn(100000),
		Name: "FairWinds",
	}
}

func (u *User) CreateStatement() {

	Statement := &Statements{
		Id:                rand.Intn(10000),
		UID:               u.Id,
		TransactionAmount: 200,
		TransactionDate:   time.Now(),
	}

	u.CheckingBalance += Statement.TransactionAmount
	u.BankStatement = append(u.BankStatement, Statement)

}

func main() {

	router := mux.NewRouter()

	bank := CreateBank()

	router.Use(loggingMiddleware)
	router.HandleFunc("/bank", bank.Details).Methods("GET")
	router.HandleFunc("/create/user", bank.CreateUser).Methods("POST")
	router.HandleFunc("/user", bank.PrintUser).Methods("GET")
	router.HandleFunc("/user/statement/{name}", bank.ViewStatement).Methods("GET")
	router.HandleFunc("/user/addTransaction/{name}", bank.AddTransaction).Methods("POST")

	http.ListenAndServe(":4000", router)

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
