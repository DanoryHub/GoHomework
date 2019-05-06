package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"repository"
)


var storage = repository.PSQLStub{}

func getUsers(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	users, err := storage.GetAllUsers()
	if err != nil{
		log.Fatal(err.Error())
	}
	_ = json.NewEncoder(writer).Encode(users)
}

func getUser(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	user, err := storage.GetUserByID(params["user_id"])
	if err != nil{
		log.Fatal(err.Error())
	}
	_ = json.NewEncoder(writer).Encode(user)
}

func getUserAccounts(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	accounts, err := storage.GetAllAccounts(params["user_id"])
	if err != nil{
		log.Fatal(err.Error())
	}
	_ = json.NewEncoder(writer).Encode(accounts)
}

func getUserAccount(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	account, err := storage.GetAccountByID(params["user_id"], params["account_id"])
	if err != nil{
		log.Fatal(err.Error())
	}
	_ = json.NewEncoder(writer).Encode(account)
}

func getAccountTransactions(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	transactions, err := storage.GetAllTransactions(params["user_id"], params["account_id"])
	if err != nil{
		log.Fatal(err.Error())
	}
	_ = json.NewEncoder(writer).Encode(transactions)
}

func getAccountTransaction(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	transaction, err := storage.GetTransactionByID(params["user_id"], params["account_id"], params["transaction_id"])
	if err != nil{
		log.Fatal(err.Error())
	}
	_ = json.NewEncoder(writer).Encode(transaction)
}

func deleteUser(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	err := storage.DeleteUser(params["user_id"])
	if err != nil{
		log.Fatal(err.Error())
	}
}

func deleteUserAccount(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	err := storage.DeleteAccount(params["user_id"], params["account_id"])
	if err != nil{
		log.Fatal(err.Error())
	}
}

func deleteAccountTransaction(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	err := storage.DeleteTransaction(params["user_id"], params["account_id"], params["transaction_id"])
	if err != nil{
		log.Fatal(err.Error())
	}
}

func createUser(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	var user repository.User
	_ = json.NewDecoder(request.Body).Decode(&user)
	user, err := storage.CreateUser(user)
	if err != nil{
		log.Fatal(err.Error())
	}
	_ = json.NewEncoder(writer).Encode(user)
}

func createUserAccount(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	var account repository.Account
	params := mux.Vars(request)
	_ = json.NewDecoder(request.Body).Decode(&account)
	account, err := storage.CreateAccount(params["user_id"], account)
	if err != nil{
		log.Fatal(err.Error())
	}
	_ = json.NewEncoder(writer).Encode(account)

}

func createAccountTransaction(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var transaction repository.Transaction
	_ = json.NewDecoder(request.Body).Decode(&transaction)
	transaction, err := storage.CreateTransaction(params["user_id"], params["account_id"], transaction)
	if err != nil{
		log.Fatal(err.Error())
	}
	_ = json.NewEncoder(writer).Encode(transaction)
}

func updateUser(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var user repository.User
	_ = json.NewDecoder(request.Body).Decode(&user)
	user, err := storage.UpdateUser(params["user_id"], user)
	if err != nil{
		log.Fatal(err.Error())
	}
	_ = json.NewEncoder(writer).Encode(user)
}

func updateUserAccount(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var account repository.Account
	_ = json.NewDecoder(request.Body).Decode(&account)
	account, err := storage.UpdateAccount(params["user_id"], params["account_id"], account)
	if err != nil{
		log.Fatal(err.Error())
	}
	_ = json.NewEncoder(writer).Encode(account)
}

func updateAccountTransaction(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	var transaction repository.Transaction
	_ = json.NewDecoder(request.Body).Decode(&transaction)
	transaction, err := storage.UpdateTransaction(params["user_id"], params["account_id"], params["transaction_id"], transaction)
	if err != nil{
		log.Fatal(err.Error())
	}
	_ = json.NewEncoder(writer).Encode(transaction)
}

func main(){
	router := mux.NewRouter()
	UsersRouter := router.PathPrefix("/users").Subrouter()
	AccountsRouter := router.PathPrefix("/users/{user_id}/accounts").Subrouter()
	TransactionsRouter := router.PathPrefix("/users/{user_id}/accounts/{account_id}/transactions").Subrouter()
	UsersRouter.HandleFunc("/", getUsers).Methods("GET")
	UsersRouter.HandleFunc("/", createUser).Methods("POST")
	UsersRouter.HandleFunc("/{user_id}", getUser).Methods("GET")
	UsersRouter.HandleFunc("/{user_id}", deleteUser).Methods("DELETE")
	UsersRouter.HandleFunc("/{user_id}", updateUser).Methods("PUT")
	AccountsRouter.HandleFunc("/", getUserAccounts).Methods("GET")
	AccountsRouter.HandleFunc("/", createUserAccount).Methods("POST")
	AccountsRouter.HandleFunc("/{account_id}", getUserAccount).Methods("GET")
	AccountsRouter.HandleFunc("/{account_id}", deleteUserAccount).Methods("DELETE")
	AccountsRouter.HandleFunc("/{account_id}", updateUserAccount).Methods("PUT")
	TransactionsRouter.HandleFunc("/", getAccountTransactions).Methods("GET")
	TransactionsRouter.HandleFunc("/", createAccountTransaction).Methods("POST")
	TransactionsRouter.HandleFunc("/{transaction_id}", getAccountTransaction).Methods("GET")
	TransactionsRouter.HandleFunc("/{transaction_id}", deleteAccountTransaction).Methods("DELETE")
	TransactionsRouter.HandleFunc("/{transaction_id}", updateAccountTransaction).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8001", router))
}
