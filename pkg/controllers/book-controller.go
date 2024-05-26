package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/yakuznecov/store-book/pkg/models"
)

var NewBook models.Book

func GetBook (w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks() // собираем все книги
	res, _ := json.Marshal(newBooks) // преобразуем в JSON
	w.Header().Set("Content-Type", "pkglication/json") // устанавливаем тип данных
	w.WriteHeader(http.StatusOK) // 200
	w.Write(res) // возвращаем книги, ответ
}