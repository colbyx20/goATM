package main

import (
	"html/template"
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

// func (u *User) CreateStatement() {

// 	Statement := &Statements{
// 		Id:                rand.Intn(10000),
// 		UID:               u.Id,
// 		TransactionAmount: 200,
// 		TransactionDate:   time.Now(),
// 	}

// 	u.CheckingBalance += Statement.TransactionAmount
// 	u.BankStatement = append(u.BankStatement, Statement)

// }

func main() {

	router := mux.NewRouter()
	bank := CreateBank()

	bank.Users["colby"] = &User{
		Id:         1,
		FirstName:  "colby",
		LastName:   "berger",
		BankNumber: 222,
		CreatedAt:  time.Now(),
	}

	var err error
	// tmpl := template.Must(template.ParseFiles("static/index.html"))
	indexTemplate = template.Must(template.ParseFiles("static/index.html"))
	userTemplate, err = template.ParseFiles("static/user.html")
	if err != nil {
		log.Fatal("Error parsing HTML template:", err)
	}

	// Page Render
	router.HandleFunc("/", IndexHandler).Methods("GET")
	router.HandleFunc("/user/login", LoggedInHandler).Methods("GET")

	// API Calls
	router.Use(loggingMiddleware)
	router.HandleFunc("/bank", bank.Details).Methods("GET")
	router.HandleFunc("/create/user", bank.CreateUser).Methods("POST")
	router.HandleFunc("/user/login", bank.PrintUser).Methods("POST")
	router.HandleFunc("/user/statement/{name}", bank.ViewStatement).Methods("GET")
	router.HandleFunc("/user/deposite/checking/{name}", bank.DepositeMoneyChecking).Methods("POST")
	router.HandleFunc("/user/deposite/savings/{name}", bank.DepositeMoneySavings).Methods("POST")
	router.HandleFunc("/user/withdraw/checking/{name}", bank.WithdrawMoneyChecking).Methods("POST")
	router.HandleFunc("/user/withdraw/savings/{name}", bank.WithdrawMoneySavings).Methods("GET")
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
