package main

import (
	"fmt"
	"github.com/izzettinozbektas/golang-api/cmd/routes"
	"github.com/izzettinozbektas/golang-api/internal/driver"
	"github.com/izzettinozbektas/golang-api/internal/handlers"
	"github.com/izzettinozbektas/golang-api/internal/helpers"
	"log"
	"net/http"
	"time"
)

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	srv := &http.Server{
		Handler:      routes.Routes(),
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start Server
	go func() {
		log.Println("Starting Server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful Shutdown
	helpers.WaitForShutdown(srv)
}
func run() (*driver.DB, error) {

	//geçici olarak kullanım,
	dbName := "golang-db"
	dbUname := "golang"
	dbPass := "golangpass"
	dbHost := "app-mysql" // mysql container name olmalı

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUname, dbPass, dbHost, dbName)
	db, err := driver.ConnectSQL(connectionString)
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}

	repo := handlers.NewRepo(db)
	handlers.NewHandlers(repo)

	return db, nil

}
