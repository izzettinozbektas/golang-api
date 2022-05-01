package routes

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/izzettinozbektas/golang-api/internal/driver"
	"github.com/izzettinozbektas/golang-api/internal/handlers"
	"github.com/izzettinozbektas/golang-api/internal/helpers"
	"net/http"
	"os"
)

var client = helpers.ConnetToRedis

func Routes() http.Handler {
	//dbName := os.Getenv("DB_NAME")
	//dbPass := os.Getenv("DB_PASS")
	//dbHost := os.Getenv("DB_HOST")
	//dbPort := os.Getenv("DB_PORT")
	// geçici olarak basit kullanım
	dbName := "golang-db"
	dbUname := "golang"
	dbPass := "golangpass"
	dbHost := "app-mysql" // mysql container name olmalı

	println("this is db", dbName, dbHost, dbPass)

	connection, err := driver.ConnectSQL(dbHost, dbUname, dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	mux := chi.NewRouter()

	mux.Get("/", handlers.Home)
	mux.Get("/redis", handlers.QuoteOfTheDayHandler())
	mux.Post("/user", handlers.NewPostHandler(connection).UserCreate)
	mux.Get("/users", handlers.NewPostHandler(connection).Users)
	mux.Post("/user/{id}", handlers.NewPostHandler(connection).UserUpdate)

	return mux
}
