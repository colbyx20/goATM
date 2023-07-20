package main

import (
	"net/http"
	"time"
)

type Banker interface {
	PrintUser(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	Details(w http.ResponseWriter, r *http.Request)
	DepositeMoney(w http.ResponseWriter, r *http.Request)
	WithdrawMoney(w http.ResponseWriter, r *http.Request)
	ViewStatement(w http.ResponseWriter, r *http.Request)
}

type TellerInf interface {
	WithdrawMoney(w http.ResponseWriter, r *http.Request)
	DepositeMoney(w http.ResponseWriter, r *http.Request)
	CheckBalance(w http.ResponseWriter, r *http.Request)
	TransferMoney(w http.ResponseWriter, r *http.Request)
}

type Bank struct {
	Id   int    `json:"id"`
	Name string `json:"bankname"`
	// Users  []*User         `json:"users"`
	Users  map[string]*User `json:"users"`
	Teller map[int]*Teller  `json:"teller"`
}

type User struct {
	Id              int           `json:"id"`
	FirstName       string        `json:"firstname"`
	LastName        string        `json:"lastname"`
	BankNumber      int           `json:"banknumber"`
	CheckingBalance float32       `json:"checkingbalance"`
	SavingsBalance  float32       `json:"savingsbalance"`
	CreatedAt       time.Time     `json:"createdat"`
	BankStatement   []*Statements `json:"bankstatement"`
}

type Statements struct {
	Id                int       `json:"id"`
	UID               int       `json:"uid"`
	TransactionAmount float32   `json:"transactionamount"`
	TransactionDate   time.Time `json:"transactiondate"`
}

type Teller struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	InUse bool   `json:"inUse"`
}
