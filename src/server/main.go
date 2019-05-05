package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"repository"
)


var storage repository.UserRepository

func getUsers(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(repository_interface.Users)
}

func getUser(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range repository_interface.Users {
		if item.ID == params["user_id"]{
			_ = json.NewEncoder(writer).Encode(item)
			return
		}
	}
	_ = json.NewEncoder(writer).Encode(&repository_interface.User{})
}

func getUserAccounts(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range repository_interface.Users {
		if item.ID == params["user_id"]{
			_ = json.NewEncoder(writer).Encode(item.Accounts)
			return
		}
	}
}

func getUserAccount(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range repository.Users {
		if item.ID == params["user_id"]{
			for _, account := range item.Accounts{
				if account.ID == params["account_id"]{
					_ = json.NewEncoder(writer).Encode(account)
					return
				}
			}
		}
	}
	_ = json.NewEncoder(writer).Encode(&repository.Account{})
}

func deleteUser(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, user := range repository.Users {
		if user.ID == params["user_id"]{
			repository.Users = append(repository.Users[:index], repository.Users[index+1:]...)
			break
		}
	}
	_ = json.NewEncoder(writer).Encode(repository.Users)
}

func deleteUserAccount(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for usr_index, user := range repository.Users {
		if user.ID == params["user_id"]{
			for index, account := range user.Accounts{
				if account.ID == params["account_id"]{
					repository.Users[usr_index].Accounts = append(repository.Users[usr_index].Accounts[:index],repository.Users[usr_index].Accounts[index+1:]...)
					break
				}
			}
			break
		}
	}
	_ = json.NewEncoder(writer).Encode(repository.Users)
}

func createUser(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	var user repository.User
	_ = json.NewDecoder(request.Body).Decode(&user)
	user.ID = strconv.Itoa(rand.Intn(10000000000))
	repository.Users = append(repository.Users, user)
	_ = json.NewEncoder(writer).Encode(user)
}

func createUserAccount(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	var account repository.Account
	params := mux.Vars(request)
	_ = json.NewDecoder(request.Body).Decode(&account)
	for index, user := range repository.Users{
		if user.ID == params["user_id"]{
			repository.Users[index].Accounts = append(repository.Users[index].Accounts, account)
			break
		}
	}
	_ = json.NewEncoder(writer).Encode(account)
}

func updateUser(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range repository.Users{
		if item.ID == params["user_id"]{
			repository.Users = append(repository.Users[:index], repository.Users[index+1:]...)
			var user repository.User
			_ = json.NewDecoder(request.Body).Decode(&user)
			user.ID = params["user_id"]
			repository.Users = append(repository.Users, user)
			_ = json.NewEncoder(writer).Encode(user)
			return
		}
	}
	_ = json.NewEncoder(writer).Encode(repository.Users)
}

func updateUserAccount(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for usr_index, item := range repository.Users{
		if item.ID == params["user_id"]{
			for index, account := range item.Accounts{
				if account.ID == params["account_id"]{
					repository.Users[usr_index].Accounts = append(repository.Users[usr_index].Accounts[:index], repository.Users[usr_index].Accounts[index+1:]...)
					var account repository.Account
					_ = json.NewDecoder(request.Body).Decode(&account)
					account.ID = params["account_id"]
					repository.Users[usr_index].Accounts = append(repository.Users[usr_index].Accounts, account)
					_ = json.NewEncoder(writer).Encode(account)
					return
				}
			}
		}
	}
	_ = json.NewEncoder(writer).Encode(repository.Users)
}

func main(){
	account1 := repository.Account{ID:"1",UserID: "1",Balance:"3800",Currency:"dollars"}
	account2 := repository.Account{ID:"2",UserID: "1",Balance:"3900",Currency:"dollars"}
	repository.Users = append(repository.Users, repository.User{ID:"1",Firstname:"Ivan",Lastname:"Metelsky",Email:"",SMS:""})
	repository.Users = append(repository.Users, repository.User{ID:"2",Firstname:"Ivan",Lastname:"Ivanov",Email:"",SMS:"+37537537537"})
	repository.Users[0].Accounts = append(repository.Users[0].Accounts, account1)
	repository.Users[0].Accounts = append(repository.Users[0].Accounts, account2)
	repository.Users[1].Accounts = append(repository.Users[1].Accounts, account1)
	repository.Users[1].Accounts = append(repository.Users[1].Accounts, account2)
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