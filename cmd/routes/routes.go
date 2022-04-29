package routes

import (
	"github.com/go-chi/chi"
	"github.com/izzettinozbektas/golang-api/internal/handlers"
	"github.com/izzettinozbektas/golang-api/internal/helpers"
	"net/http"
)

var client = helpers.ConnetToRedis

func Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Home)
	mux.Get("/redis", handlers.QuoteOfTheDayHandler())

	return mux
}
