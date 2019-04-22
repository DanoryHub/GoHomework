package main

import(
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var books []Book

func GetBooks(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(books)
}

func GetBook(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range books{
		if item.ID == params["id"]{
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
	json.NewEncoder(writer)
}

func CreateBook(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type","application/json")
	var book Book
	_ = json.NewDecoder(request.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	books = append(books, book)
	json.NewEncoder(writer).Encode(book)
}

func UpdateBook(writer http.ResponseWriter, request *http.Request){
	writer.Header().Set("Content-Type","application/json")
	params := mux.Vars(request)
	for index, item := range books{
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(request.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(writer).Encode(book)
			return
		}
	}
	json.NewEncoder(writer).Encode(books)
}

func DeleteBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
		json.NewEncoder(writer).Encode(books)
	}
}

func main(){
	r := mux.NewRouter()
	books = append(books, Book{ID: "1", Title: "Война и Мир", Author: &Author{Firstname: "Лев", Lastname: "Толстой"}})
	books = append(books, Book{ID: "2", Title: "Преступление и наказание", Author: &Author{Firstname: "Фёдор", Lastname: "Достоевский"}})
	r.HandleFunc("/books", GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", GetBook).Methods("GET")
	r.HandleFunc("/books", CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
