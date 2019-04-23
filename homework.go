package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	_ "log"
	"math/rand"
	_ "math/rand"
	"net/http"
	"strconv"
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

type UserRepository interface {
	GetAll() ([]User, error)
	FindByID(id string) (User, error)
	UpdateUser(id string, user User) (User, error)

}

var storage UserRepository

func getUsers(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(users)
}

func getUser(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for _, item := range users{
		if item.ID == params["user_id"]{
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

func deleteUserAccount(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for usr_index, user := range users{
		if user.ID == params["user_id"]{
			for index, account := range user.Accounts{
				if account.ID == params["account_id"]{
					users[usr_index].Accounts = append(users[usr_index].Accounts[:index],users[usr_index].Accounts[index+1:]...)
					break
				}
			}
			break
		}
	}
	_ = json.NewEncoder(writer).Encode(users)
}

func createUser(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(request.Body).Decode(&user)
	user.ID = strconv.Itoa(rand.Intn(10000000000))
	users = append(users, user)
	_ = json.NewEncoder(writer).Encode(user)
}

func createUserAccount(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	var account Account
	params := mux.Vars(request)
	_ = json.NewDecoder(request.Body).Decode(&account)
	for index, user := range users{
		if user.ID == params["user_id"]{
			users[index].Accounts = append(users[index].Accounts, account)
			break
		}
	}
	_ = json.NewEncoder(writer).Encode(account)
}

func updateUser(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range users{
		if item.ID == params["user_id"]{
			users = append(users[:index], users[index+1:]...)
			var user User
			_ = json.NewDecoder(request.Body).Decode(&user)
			user.ID = params["user_id"]
			users = append(users, user)
			_ = json.NewEncoder(writer).Encode(user)
			return
		}
	}
	_ = json.NewEncoder(writer).Encode(users)
}

func updateUserAccount(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for usr_index, item := range users{
		if item.ID == params["user_id"]{
			for index, account := range item.Accounts{
				if account.ID == params["account_id"]{
					users[usr_index].Accounts = append(users[usr_index].Accounts[:index], users[usr_index].Accounts[index+1:]...)
					var account Account
					_ = json.NewDecoder(request.Body).Decode(&account)
					account.ID = params["account_id"]
					users[usr_index].Accounts = append(users[usr_index].Accounts, account)
					_ = json.NewEncoder(writer).Encode(account)
					return
				}
			}
		}
	}
	_ = json.NewEncoder(writer).Encode(users)
}

func main(){
	account1 := Account{ID:"1",UserID: "1",Balance:"3800",Currency:"dollars"}
	account2 := Account{ID:"2",UserID: "1",Balance:"3900",Currency:"dollars"}
	users = append(users, User{ID:"1",Firstname:"Ivan",Lastname:"Metelsky",Email:"",SMS:""})
	users = append(users, User{ID:"2",Firstname:"Ivan",Lastname:"Ivanov",Email:"",SMS:"+37537537537"})
	users[0].Accounts = append(users[0].Accounts, account1)
	users[0].Accounts = append(users[0].Accounts, account2)
	users[1].Accounts = append(users[1].Accounts, account1)
	users[1].Accounts = append(users[1].Accounts, account2)
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users/{user_id}", getUser).Methods("GET")
	router.HandleFunc("/users/{user_id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/users/{user_id}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{user_id}/accounts", getUserAccounts).Methods("GET")
	router.HandleFunc("/users/{user_id}/accounts", createUserAccount).Methods("POST")
	router.HandleFunc("/users/{user_id}/accounts/{account_id}", getUserAccount).Methods("GET")
	router.HandleFunc("/users/{user_id}/accounts/{account_id}", deleteUserAccount).Methods("DELETE")
	router.HandleFunc("/users/{user_id}/accounts/{account_id}", updateUserAccount).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", router))
}