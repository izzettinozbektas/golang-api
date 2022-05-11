package routes

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/izzettinozbektas/golang-api/cmd/middleware"
	"github.com/izzettinozbektas/golang-api/internal/driver"
	"github.com/izzettinozbektas/golang-api/internal/handlers"
	"github.com/izzettinozbektas/golang-api/internal/helpers"
	"net/http"
	"os"
)

var client = helpers.ConnetToRedis

func Routes() http.Handler {
	connection, err := driver.ConnectSQL()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	mux := chi.NewRouter()

	mux.Get("/", handlers.Home)
	mux.Get("/redis", handlers.Redis)
	//login
	mux.Post("/login", handlers.ConnectionToDB(connection).Login)
	//user singup
	mux.Post("/user", handlers.ConnectionToDB(connection).UserCreate)

	mux.With(middleware.Auth).Group(func(r chi.Router) {
		r.Get("/users", handlers.ConnectionToDB(connection).Users)
		r.Put("/user/{id}", handlers.ConnectionToDB(connection).UserUpdate)
		r.Get("/user/{id}", handlers.ConnectionToDB(connection).User)
		r.Delete("/user/{id}", handlers.ConnectionToDB(connection).UserDelete)
		return
	})

	return mux
}
