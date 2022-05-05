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
	mux.Get("/redis", handlers.Redis)
	//user Route
	mux.Post("/user", handlers.ConnectionToDB(connection).UserCreate)
	mux.Get("/users", handlers.ConnectionToDB(connection).Users)
	mux.Put("/user/{id}", handlers.ConnectionToDB(connection).UserUpdate)
	mux.Get("/user/{id}", handlers.ConnectionToDB(connection).User)
	mux.Delete("/user/{id}", handlers.ConnectionToDB(connection).UserDelete)

	return mux
}
