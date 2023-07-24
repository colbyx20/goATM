package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const (
	Checking = "checking"
	Savings  = "savings"
)

var indexTemplate *template.Template
var userTemplate *template.Template

// Function to render the HTML template
func renderHTMLTemplate(w http.ResponseWriter, tmpl *template.Template, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, data)
}

// Function to render the "LoggedInHandler" page with the updated data
func renderLoggedInHandler(w http.ResponseWriter, r *http.Request, b *Bank, user *User) {
	data := struct {
		User *User // User data
	}{
		User: user,
	}

	renderHTMLTemplate(w, userTemplate, data)
}

// Handler for the index page (GET request)
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// You can pass any initial data to the template here (if needed)
	data := struct {
		User *User // Initial data (optional)
	}{
		User: nil, // Initial user data (set to nil initially)
	}

	// Render the HTML template with the initial data
	renderHTMLTemplate(w, indexTemplate, data)
}

// Handler for the index page (GET request)
func LoggedInHandler(w http.ResponseWriter, r *http.Request) {
	// You can pass any initial data to the template here (if needed)
	data := struct {
		User *User // Initial data (optional)
	}{
		User: nil, // Initial user data (set to nil initially)
	}

	// Render the HTML template with the initial data
	renderHTMLTemplate(w, indexTemplate, data)
}

func (b *Bank) CreateUser(w http.ResponseWriter, r *http.Request) {

	// parse data from form
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse for data", http.StatusInternalServerError)
		return
	}

	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")

	// u := new(User)
	// json.NewDecoder(r.Body).Decode(u)
	// defer r.Body.Close()

	_, ok := b.Users[firstName]

	if !ok {
		// u := &User{
		// 	Id:              rand.Intn(1000),
		// 	FirstName:       firstName,
		// 	LastName:        lastName,
		// 	BankNumber:      rand.Intn(100000000),
		// 	CheckingBalance: 0,
		// 	SavingsBalance:  0,
		// 	CreatedAt:       time.Now(),
		// }

		// b.Users[u.FirstName] = u

		// w.WriteHeader(http.StatusCreated)
		// json.NewEncoder(w).Encode(u)
		// return

		newUser := &User{
			// Assuming you have fields like Id, BankNumber, CheckingBalance, SavingsBalance, etc.
			Id:              rand.Intn(1000),
			BankNumber:      rand.Intn(100000000),
			CheckingBalance: 0,
			SavingsBalance:  0,
			CreatedAt:       time.Now(),
			FirstName:       firstName,
			LastName:        lastName,
		}

		// Save the user data to the bank
		b.Users[firstName] = newUser

		// You can also pass the newly created user data to the template for display
		data := struct {
			User *User // New user data
		}{
			User: newUser,
		}

		renderHTMLTemplate(w, indexTemplate, data)

	} else {
		json.NewEncoder(w).Encode(map[string]string{"User Already Exists!": firstName})
		return
	}

}

func (b *Bank) Details(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK) // send 200
	json.NewEncoder(w).Encode(map[string]*Bank{b.Name: b})
}

func (b *Bank) DepositeMoneyChecking(w http.ResponseWriter, r *http.Request) {

	name := mux.Vars(r)["name"]

	fmt.Println("Name: ", name)
	newTransaction := new(Statements)

	err := r.ParseForm()
	if err != nil {
		log.Fatal("err :", err)
	}

	// err := r.FormValue("transactionType")
	amountStr := r.FormValue("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		fmt.Println("error")
	}

	// Decode incoming request from the r.Body
	json.NewDecoder(r.Body).Decode(newTransaction)
	defer r.Body.Close()

	newTransaction.Id = rand.Intn(10000)
	newTransaction.UID = b.Users[name].Id
	newTransaction.AccountType = Checking
	newTransaction.TransactionType = "Deposit"
	newTransaction.TransactionDate = time.Now().Format("2023-01-01")
	newTransaction.TransactionAmount = float32(amount)
	u, ok := b.Users[name]

	if ok {
		u.CheckingBalance += newTransaction.TransactionAmount
		u.BankStatement = append(u.BankStatement, newTransaction)

		data := struct {
			User       *User // New user data
			Bank       *Bank
			Statements []*Statements
		}{
			User:       u,
			Bank:       b,
			Statements: u.BankStatement,
		}

		renderHTMLTemplate(w, userTemplate, data)
		return
	} else {
		json.NewEncoder(w).Encode(map[string]string{"User Doesn't Exist: ": name})
		return
	}
}

func (b *Bank) DepositeMoneySavings(w http.ResponseWriter, r *http.Request) {

	name := mux.Vars(r)["name"]

	newTransaction := new(Statements)

	err := r.ParseForm()
	if err != nil {
		log.Fatal("err :", err)
	}

	// err := r.FormValue("transactionType")
	amountStr := r.FormValue("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		fmt.Println("error")
	}

	// Decode incoming request from the r.Body
	json.NewDecoder(r.Body).Decode(newTransaction)
	defer r.Body.Close()

	newTransaction.Id = rand.Intn(10000)
	newTransaction.UID = b.Users[name].Id
	newTransaction.AccountType = Checking
	newTransaction.TransactionType = "Deposit"
	newTransaction.TransactionDate = time.Now().Format("2023-01-01")
	newTransaction.TransactionAmount = float32(amount)

	u, ok := b.Users[name]

	if ok {
		u.SavingsBalance += newTransaction.TransactionAmount
		u.BankStatement = append(u.BankStatement, newTransaction)

		data := struct {
			User       *User // New user data
			Bank       *Bank
			Statements []*Statements
		}{
			User:       u,
			Bank:       b,
			Statements: u.BankStatement,
		}

		// indexTemplate = template.Must(template.ParseFiles("static/index.html"))
		// indexTemplate.Execute(w,data)

		renderHTMLTemplate(w, userTemplate, data)

		// json.NewEncoder(w).Encode(&newTransaction)
		return
	} else {
		json.NewEncoder(w).Encode(map[string]string{"User Doesn't Exist: ": name})
		return
	}
}

func (b *Bank) WithdrawMoneyChecking(w http.ResponseWriter, r *http.Request) {

	name := mux.Vars(r)["name"]

	err := r.ParseForm()
	if err != nil {
		log.Fatal("err :", err)
	}

	// err := r.FormValue("transactionType")
	amountStr := r.FormValue("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		fmt.Println("error")
	}
	// checking := r.FormValue("checking")

	newTransaction := new(Statements)

	// Decode incoming request from the r.Body
	json.NewDecoder(r.Body).Decode(newTransaction)
	defer r.Body.Close()

	newTransaction.Id = rand.Intn(10000)
	newTransaction.UID = b.Users[name].Id
	newTransaction.AccountType = "Checking"
	newTransaction.TransactionType = "Withdraw"
	newTransaction.TransactionDate = time.Now().Format("2023-01-01")
	newTransaction.TransactionAmount = float32(amount)

	// does user exists?
	u, ok := b.Users[name]

	if ok {
		u.CheckingBalance -= float32(amount)
		u.BankStatement = append(u.BankStatement, newTransaction)

		// json.NewEncoder(w).Encode(&newTransaction)
		data := struct {
			User       *User // New user data
			Bank       *Bank
			Statements []*Statements
		}{
			User:       u,
			Bank:       b,
			Statements: u.BankStatement,
		}

		renderHTMLTemplate(w, userTemplate, data)

		return
	} else {
		json.NewEncoder(w).Encode(map[string]string{"User Doesn't Exist: ": name})
		return
	}
}

func (b *Bank) WithdrawMoneySavings(w http.ResponseWriter, r *http.Request) {

	name := mux.Vars(r)["name"]

	newTransaction := new(Statements)

	amountStr := r.FormValue("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		fmt.Println("error")
	}

	// Decode incoming request from the r.Body
	json.NewDecoder(r.Body).Decode(newTransaction)
	defer r.Body.Close()

	newTransaction.Id = rand.Intn(10000)
	newTransaction.UID = b.Users[name].Id
	newTransaction.AccountType = Savings
	newTransaction.TransactionType = "Withdraw"
	newTransaction.TransactionDate = time.Now().Format("2023-01-01")
	newTransaction.TransactionAmount = float32(amount)

	// does user exists?
	u, ok := b.Users[name]

	if ok {
		u.SavingsBalance -= float32(amount)
		u.BankStatement = append(u.BankStatement, newTransaction)

		data := struct {
			User       *User // New user data
			Bank       *Bank
			Statements []*Statements
		}{
			User:       u,
			Bank:       b,
			Statements: u.BankStatement,
		}

		// indexTemplate = template.Must(template.ParseFiles("static/index.html"))
		// indexTemplate.Execute(w,data)

		renderHTMLTemplate(w, userTemplate, data)

		// json.NewEncoder(w).Encode(&newTransaction)
		return
	} else {
		json.NewEncoder(w).Encode(map[string]string{"User Doesn't Exist: ": name})
		return
	}
}

func (b *Bank) PrintUser(w http.ResponseWriter, r *http.Request) {

	// name := mux.Vars(r)["name"]
	// json.NewDecoder(r.Body).Decode(&name)

	// grab data from form

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Error on form")
	}

	name := r.FormValue("firstName")
	user, ok := b.Users[name]

	if ok {

		data := struct {
			User       *User // New user data
			Bank       *Bank
			Statements []*Statements
		}{
			User:       user,
			Bank:       b,
			Statements: user.BankStatement,
		}

		// indexTemplate = template.Must(template.ParseFiles("static/index.html"))
		// indexTemplate.Execute(w,data)

		renderHTMLTemplate(w, userTemplate, data)
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
