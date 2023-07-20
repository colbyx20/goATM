package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func UserExist(user map[string]*User, name string) bool {

	_, ok := user[name]
	if ok {
		return false
	}
	return true

}

func (b *Bank) CreateUser(w http.ResponseWriter, r *http.Request) {

	u := new(User)
	json.NewDecoder(r.Body).Decode(u)
	defer r.Body.Close()

	// for _, user := range b.Users {
	// 	if user.FirstName == u.FirstName {
	// 		json.NewEncoder(w).Encode(map[string]string{"User Already Exists! ": u.FirstName})
	// 		return
	// 	}
	// }

	ok := UserExist(b.Users, u.FirstName)
	if ok {

		u.Id = rand.Intn(1000)
		u.BankNumber = rand.Intn(100000000)
		u.CheckingBalance = 0
		u.SavingsBalance = 0
		u.CreatedAt = time.Now()
		b.Users[u.FirstName] = u

		// b.Users = append(b.Users, u)
		json.NewEncoder(w).Encode(b.Users)

	} else {
		json.NewEncoder(w).Encode(map[string]string{"User Already Exists!": u.FirstName})
	}

}

func (b *Bank) Details(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK) // send 200
	json.NewEncoder(w).Encode(map[string]*Bank{b.Name: b})
}

func (b *Bank) DepositeMoney(w http.ResponseWriter, r *http.Request) {

	name := mux.Vars(r)["name"]

	newTransaction := new(Statements)

	// Decode incoming request from the r.Body
	json.NewDecoder(r.Body).Decode(newTransaction)
	defer r.Body.Close()

	newTransaction.Id = rand.Intn(10000)
	newTransaction.UID = b.Id
	newTransaction.TransactionDate = time.Now()

	fmt.Println(name)
	// does user exists?
	ok := UserExist(b.Users, name)
	fmt.Println(b.Users)
	fmt.Println(ok)
	if !ok {
		b.Users[name].CheckingBalance += newTransaction.TransactionAmount
		b.Users[name].BankStatement = append(b.Users[name].BankStatement, newTransaction)
		json.NewEncoder(w).Encode(&newTransaction)

	} else {
		json.NewEncoder(w).Encode(map[string]string{"User Doesn't Exist: ": name})
	}

	// // find the user with name name
	// for _, user := range b.Users {
	// 	if user.FirstName == name {
	// 		user.CheckingBalance += newTransaction.TransactionAmount
	// 		json.NewEncoder(w).Encode(&newTransaction)
	// 	} else {
	// 		json.NewEncoder(w).Encode(map[string]string{"User Doesn't Exist: ": name})
	// 	}
	// }
}

func (b *Bank) PrintUser(w http.ResponseWriter, r *http.Request) {

	// w.WriteHeader(http.StatusOK) // send 200
	// json.NewEncoder(w).Encode(map[string][]*User{"test": b.Users})
}

func (b *Bank) ViewStatement(w http.ResponseWriter, r *http.Request) {

	// grab a user
	name := mux.Vars(r)["name"]

	var currUser *User
	// find user
	for idx, user := range b.Users {
		if user.FirstName == name {
			currUser = b.Users[idx]
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string][]*Statements{currUser.FirstName + " " + currUser.LastName: currUser.BankStatement})
}

// func (b *Bank) DepositeMoney(w http.ResponseWriter, r *http.Request) {

// 	// var id int
// 	// json.NewDecoder(r.Body).Decode(id)
// 	// fmt.Println(id)
// 	// json.NewEncoder(w).Encode(b)

// }

func (b *Bank) WithdrawMoney(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK) // send 200
	json.NewEncoder(w).Encode("HI;lkj;LKj")
}

func (b *Bank) CheckBalance(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK) // send 200
	json.NewEncoder(w).Encode("HI;lkj;LKj")
}

func (b *Bank) TransferMoney(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK) // send 200
	json.NewEncoder(w).Encode("HI;lkj;LKj")
}
