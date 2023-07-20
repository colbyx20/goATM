package main

import "time"

type Bank struct {
	Id    int     `json:"id"`
	Name  string  `json:"bankname"`
	Users []*User `json:"users"`
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
