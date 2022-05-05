package main

import (
	"github.com/izzettinozbektas/golang-api/cmd/routes"
	"github.com/izzettinozbektas/golang-api/internal/helpers"
	"log"
	"net/http"
	"time"
)

func main() {
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
