package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (t *Teller) DepositeMoney(w http.ResponseWriter, r *http.Request) {

	fmt.Println(t)

}

func (t *Teller) WithdrawMoney(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK) // send 200
	json.NewEncoder(w).Encode("HI;lkj;LKj")
}

func (t *Teller) CheckBalance(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK) // send 200
	json.NewEncoder(w).Encode("HI;lkj;LKj")
}

func (t *Teller) TransferMoney(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK) // send 200
	json.NewEncoder(w).Encode("HI;lkj;LKj")
}
