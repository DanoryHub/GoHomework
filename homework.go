package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	_ "log"
	_ "math/rand"
	"net/http"
	_ "strconv"
)

type User struct {
	ID string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	SMS string `json:"sms"`
	Accounts []Account `json:"accounts"`
}

type Account struct{
	ID string `json:"id"`
	UserID string `json:"user_id"`
	Balance string `json:"balance"`
	Currency string `json:"currency"`
}

var users []User

func getUsers(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(users)
}

func getUser(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range users{
		if item.ID == params["id"]{
			_ = json.NewEncoder(writer).Encode(item)
			return
		}
	}
	_ = json.NewEncoder(writer).Encode(&User{})
}

func getUserAccounts(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range users{
		if item.ID == params["user_id"]{
			_ = json.NewEncoder(writer).Encode(item.Accounts)
			return
		}
	}
}

func getUserAccount(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range users{
		if item.ID == params["user_id"]{
			for _, account := range item.Accounts{
				if account.ID == params["account_id"]{
					_ = json.NewEncoder(writer).Encode(account)
					return
				}
			}
		}
	}
	_ = json.NewEncoder(writer).Encode(&Account{})
}

func deleteUser(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, user := range users{
		if user.ID == params["user_id"]{
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	_ = json.NewEncoder(writer).Encode(users)
}

func main(){
	router := mux.NewRouter()
	account1 := Account{ID:"1",UserID: "1",Balance:"3800",Currency:"dollars"}
	account2 := Account{ID:"2",UserID: "1",Balance:"3900",Currency:"dollars"}
	users = append(users, User{ID:"1",Firstname:"Ivan",Lastname:"Metelsky",Email:"",SMS:""})
	users = append(users, User{ID:"2",Firstname:"Ivan",Lastname:"Ivanov",Email:"",SMS:"+37537537537"})
	users[0].Accounts = append(users[0].Accounts, account1)
	users[0].Accounts = append(users[0].Accounts, account2)
	users[1].Accounts = append(users[1].Accounts, account1)
	users[1].Accounts = append(users[1].Accounts, account2)
	router.HandleFunc("/users", getUsers).Methods("Get")
	router.HandleFunc("/users/{id}", getUser).Methods("Get")
	router.HandleFunc("/users/{user_id}/accounts", getUserAccounts).Methods("Get")
	router.HandleFunc("/users/{user_id}/accounts/{account_id}", getUserAccount).Methods("Get")
	router.HandleFunc("/users/{user_id}", deleteUser).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8000", router))
}