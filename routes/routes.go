package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
	"github.com/izzettinozbektas/golang-api/internal/handlers"
	"net/http"
)

func Routes(client *redis.Client) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Home)
	mux.Get("/redis", handlers.QuoteOfTheDayHandler(client))

	return mux
}
