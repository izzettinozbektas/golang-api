package routes

import (
	"github.com/go-chi/chi"
	"github.com/izzettinozbektas/golang-api/cmd/middleware"
	"github.com/izzettinozbektas/golang-api/internal/handlers"
	"github.com/izzettinozbektas/golang-api/internal/helpers"
	"net/http"
)

var client = helpers.ConnetToRedis

func Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Home)
	mux.Get("/redis", handlers.Redis)
	//login
	mux.Post("/login", handlers.Repo.Login)
	//user singup
	mux.Post("/user", handlers.Repo.UserCreate)

	mux.With(middleware.Auth).Group(func(r chi.Router) {
		r.Get("/getUserData", handlers.Repo.GetUserInformation)
		r.Get("/users", handlers.Repo.Users)
		r.Put("/user/{id}", handlers.Repo.UserUpdate)
		r.Get("/user/{id}", handlers.Repo.User)
		r.Delete("/user/{id}", handlers.Repo.UserDelete)
		return
	})

	return mux
}
