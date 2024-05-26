package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yakuznecov/store-book/pkg/routes" // импорт маршрутов
)

func main() {
	r := mux.NewRouter() // инициализация маршрутизатора
	routes.RegisterBookStoreRoutes(r) // передаем роутер
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}