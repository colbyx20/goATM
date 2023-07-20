package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// constructor - creat user
func CreateBank() *Bank {
	return &Bank{
		Id:    rand.Intn(100000),
		Name:  "FairWinds",
		Users: make(map[string]*User),
	}
}

func CreateTeller() *Teller {
	return &Teller{
		Id:   1,
		Name: "Teller",
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
	// bank.Teller = make(map[int]*Teller)
	// t1 := CreateTeller()
	// bank.Teller[t1.Id] = t1

	// fmt.Println(*bank.Teller[1])

	router.Use(loggingMiddleware)
	router.HandleFunc("/bank", bank.Details).Methods("GET")
	router.HandleFunc("/create/user", bank.CreateUser).Methods("POST")
	router.HandleFunc("/user/{name}", bank.PrintUser).Methods("GET")
	router.HandleFunc("/user/statement/{name}", bank.ViewStatement).Methods("GET")
	router.HandleFunc("/user/deposite/{name}", bank.DepositeMoney).Methods("POST")
	router.HandleFunc("/user/withdraw/{name}", bank.WithdrawMoney).Methods("POST")
	router.HandleFunc("/user/balance/{name}", bank.CheckBalance).Methods("GET")

	// router.HandleFunc("/bank/teller/deposit", bank.DepositeMoney).Methods("POST")
	// router.HandleFunc("/bank/teller/withdraw", bank.WithdrawMoney).Methods("POST")
	// router.HandleFunc("/bank/teller/balance", bank.CheckBalance).Methods("GET")
	// router.HandleFunc("/bank/teller/transfer", bank.TransferMoney).Methods("POST")

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
