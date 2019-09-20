package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// Account is used to hold the account details
type Account struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phonenumber"`
}

// acc is used to hold the account details
var acc []Account

// example handler
func createHandler(w http.ResponseWriter, r *http.Request) {

	// read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var a Account
	// Unmarshal the body into variable a
	err = json.Unmarshal(body, &a)
	if err != nil {
		panic(err)
	}
	// store the values
	acc = append(acc, a)

	// write the response back
	w.Write([]byte("Account created successfully"))
}
func getAccountDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// if no accounts, respond with not found status
	if len(acc) == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	// encode data
	accJSON, err := json.Marshal(&acc)
	if err != nil {
		// if there is any error, write the
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
	}
	// write the response
	w.WriteHeader(http.StatusOK)
	w.Write(accJSON)
}
func updateAccountHandler(w http.ResponseWriter, r *http.Request) {
	//read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}
	var a Account
	var ac []Account
	//Unmarshal the body into variable a
	err = json.Unmarshal(body, &a)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}

	for i := range acc {

		if a.Name != "" {
			acc[i].Name = a.Name
		}
		if a.Email != "" {
			acc[i].Email = a.Email
		}
		if a.PhoneNumber != "" {
			acc[i].PhoneNumber = a.PhoneNumber
		}

		ac = append(ac, acc[i])
	}

	//update account
	acc = ac
	w.Write([]byte("Account updated successfully"))
}

func deleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	// read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}

	var a Account
	var ac []Account
	var j int
	//Unmarshal the body into variable a
	err = json.Unmarshal(body, &a)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}

	for i := range acc {
		if acc[i].ID == a.ID {
			j++
		} else {
			ac = append(ac, acc[i])
		}
	}
	//delete account
	acc = ac
	if j > 0 {
		w.Write([]byte("Account deleted successfully"))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Account not found"))
	}

}

// build the router
func makeRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/create", createHandler)
	r.HandleFunc("/get", getAccountDetailsHandler)
	r.HandleFunc("/update", updateAccountHandler)
	r.HandleFunc("/delete", deleteAccountHandler)
	http.ListenAndServe(":8081", r)
	return r
}

func main() {
	// initialize the router
	makeRouter()
}
