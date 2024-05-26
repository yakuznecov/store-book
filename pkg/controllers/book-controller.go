package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yakuznecov/store-book/pkg/models"
	"github.com/yakuznecov/store-book/pkg/utils"
)

var NewBook models.Book

func GetBook (w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks() // собираем все книги
	res, _ := json.Marshal(newBooks) // преобразуем в JSON
	w.Header().Set("Content-Type", "pkglication/json") // устанавливаем тип данных
	w.WriteHeader(http.StatusOK) // 200
	w.Write(res) // возвращаем книги, ответ
}

func GetBookById (w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r) // получаем параметры
	bookId := vars["bookId"] // получаем ID
	ID, err := strconv.ParseInt(bookId, 0, 0) // преобразуем ID
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID) // получаем книгу
	res, _ := json.Marshal(bookDetails) // преобразуем в JSON
	w.Header().Set("Content-Type", "pkglication/json") // устанавливаем тип данных
	w.WriteHeader(http.StatusOK) // 200
	w.Write(res) // возвращаем книгу, ответ
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook) // принимаем данные
	b := CreateBook.CreateBook() // создаем книгу
	res, _ := json.Marshal(b) // преобразуем в JSON
	w.WriteHeader(http.StatusOK) // 200
	w.Write(res) // возвращаем книгу, ответ
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r) // получаем параметры
	bookId := vars["bookId"] // получаем ID
	ID, err := strconv.ParseInt(bookId, 0, 0) // преобразуем ID
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID) // удаляем книгу
	res, _ := json.Marshal(book) // преобразуем в JSON
	w.Header().Set("Content-Type", "pkglication/json") // устанавливаем тип данных
	w.WriteHeader(http.StatusOK) // 200
	w.Write(res) // возвращаем книгу, ответ
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook) // принимаем данные
	vars := mux.Vars(r) // получаем параметры
	bookId := vars["bookId"] // получаем ID
	ID, err := strconv.ParseInt(bookId, 0, 0) // преобразуем ID
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(ID) // получаем книгу
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails) // обновляем книгу
	res, _ := json.Marshal(bookDetails) // преобразуем в JSON
	w.Header().Set("Content-Type", "pkglication/json") // устанавливаем тип данных
	w.WriteHeader(http.StatusOK) // 200
	w.Write(res) // возвращаем книгу, ответ
}