package main

import (
	"encoding/json"
	"html/template"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func renderHTMLTemplate(w http.ResponseWriter, tmpl *template.Template, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, data)
}

func UserHandler(b *Bank) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)["name"]
		user, ok := b.Users[name]

		if !ok {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		data := struct {
			User *User
		}{
			User: user,
		}

		renderHTMLTemplate(w, userTemplate, data)
	}
}

func (b *Bank) CreateUser(w http.ResponseWriter, r *http.Request) {

	u := new(User)
	json.NewDecoder(r.Body).Decode(u)
	defer r.Body.Close()

	_, ok := b.Users[u.FirstName]

	if !ok {
		u.Id = rand.Intn(1000)
		u.BankNumber = rand.Intn(100000000)
		u.CheckingBalance = 0
		u.SavingsBalance = 0
		u.CreatedAt = time.Now()
		b.Users[u.FirstName] = u

		// b.Users = append(b.Users, u)
		// json.NewEncoder(w).Encode(b.Users)
		http.Redirect(w, r, "/user/"+u.FirstName, http.StatusSeeOther)
		return

	} else {
		json.NewEncoder(w).Encode(map[string]string{"User Already Exists!": u.FirstName})
		return
	}

}

func (b *Bank) Details(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK) // send 200
	json.NewEncoder(w).Encode(map[string]*Bank{b.Name: b})
}

func (b *Bank) DepositeMoneyChecking(w http.ResponseWriter, r *http.Request) {

	name := mux.Vars(r)["name"]

	newTransaction := new(Statements)

	// Decode incoming request from the r.Body
	json.NewDecoder(r.Body).Decode(newTransaction)
	defer r.Body.Close()

	newTransaction.Id = rand.Intn(10000)
	newTransaction.UID = b.Users[name].Id
	newTransaction.TransactionType = "Deposit"
	newTransaction.TransactionDate = time.Now()

	u, ok := b.Users[name]

	if ok {
		u.CheckingBalance += newTransaction.TransactionAmount
		u.BankStatement = append(u.BankStatement, newTransaction)
		json.NewEncoder(w).Encode(&newTransaction)
		return
	} else {
		json.NewEncoder(w).Encode(map[string]string{"User Doesn't Exist: ": name})
		return
	}
}

func (b *Bank) DepositeMoneySavings(w http.ResponseWriter, r *http.Request) {

	name := mux.Vars(r)["name"]

	newTransaction := new(Statements)

	// Decode incoming request from the r.Body
	json.NewDecoder(r.Body).Decode(newTransaction)
	defer r.Body.Close()

	newTransaction.Id = rand.Intn(10000)
	newTransaction.UID = b.Users[name].Id
	newTransaction.TransactionType = "Deposit"
	newTransaction.TransactionDate = time.Now()

	u, ok := b.Users[name]

	if ok {
		u.SavingsBalance += newTransaction.TransactionAmount
		u.BankStatement = append(u.BankStatement, newTransaction)
		json.NewEncoder(w).Encode(&newTransaction)
		return
	} else {
		json.NewEncoder(w).Encode(map[string]string{"User Doesn't Exist: ": name})
		return
	}
}

func (b *Bank) WithdrawMoneyChecking(w http.ResponseWriter, r *http.Request) {

	name := mux.Vars(r)["name"]

	newTransaction := new(Statements)

	// Decode incoming request from the r.Body
	json.NewDecoder(r.Body).Decode(newTransaction)
	defer r.Body.Close()

	newTransaction.Id = rand.Intn(10000)
	newTransaction.UID = b.Users[name].Id
	newTransaction.TransactionType = "Withdraw"
	newTransaction.TransactionDate = time.Now()

	// does user exists?
	u, ok := b.Users[name]

	if ok {
		u.CheckingBalance -= newTransaction.TransactionAmount
		u.BankStatement = append(u.BankStatement, newTransaction)
		json.NewEncoder(w).Encode(&newTransaction)
		return
	} else {
		json.NewEncoder(w).Encode(map[string]string{"User Doesn't Exist: ": name})
		return
	}
}

func (b *Bank) WithdrawMoneySavings(w http.ResponseWriter, r *http.Request) {

	name := mux.Vars(r)["name"]

	newTransaction := new(Statements)

	// Decode incoming request from the r.Body
	json.NewDecoder(r.Body).Decode(newTransaction)
	defer r.Body.Close()

	newTransaction.Id = rand.Intn(10000)
	newTransaction.UID = b.Users[name].Id
	newTransaction.TransactionType = "Withdraw"
	newTransaction.TransactionDate = time.Now()

	// does user exists?
	u, ok := b.Users[name]

	if ok {
		u.SavingsBalance -= newTransaction.TransactionAmount
		u.BankStatement = append(u.BankStatement, newTransaction)
		json.NewEncoder(w).Encode(&newTransaction)
		return
	} else {
		json.NewEncoder(w).Encode(map[string]string{"User Doesn't Exist: ": name})
		return
	}
}

func (b *Bank) PrintUser(w http.ResponseWriter, r *http.Request) {

	name := mux.Vars(r)["name"]
	json.NewDecoder(r.Body).Decode(&name)

	user, ok := b.Users[name]

	if ok {
		w.WriteHeader(http.StatusOK) // send 200
		json.NewEncoder(w).Encode(user)
		return
	} else {
		json.NewEncoder(w).Encode(map[string]string{"User doesn't Exist!": name})
		return
	}

}

func (b *Bank) ViewStatement(w http.ResponseWriter, r *http.Request) {

	// grab a user
	name := mux.Vars(r)["name"]
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(b.Users[name].BankStatement)
}

func (b *Bank) CheckBalance(w http.ResponseWriter, r *http.Request) {

	name := mux.Vars(r)["name"]

	user, ok := b.Users[name]

	if ok {

		json.NewEncoder(w).Encode(map[string]float32{
			"Checking Balance": user.CheckingBalance,
			"Savings Balance":  user.SavingsBalance,
		})
		return
	} else {
		json.NewEncoder(w).Encode(map[string]string{"User doesn't Exist!": name})
		return
	}

}

func (b *Bank) TransferMoney(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK) // send 200
	json.NewEncoder(w).Encode("HI;lkj;LKj")
}
