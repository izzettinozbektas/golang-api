package handlers

import (
	"github.com/go-redis/redis"
	"github.com/izzettinozbektas/golang-api/internal/helpers"
	"log"
	"net/http"
	"time"
)

// redis client connect
var client = helpers.ConnetToRedis()

// home is the handler
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome! Please hit the `/redis` API to get the quote of the day."))
}

// redis the handler
func QuoteOfTheDayHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now()
		date := currentTime.Format("2006-01-02")

		val, err := client.Get(date).Result()
		if err == redis.Nil {
			log.Println("Cache miss for date ", date)
			quoteResp, err := helpers.GetQuoteFromAPI()
			if err != nil {
				w.Write([]byte("Sorry! We could not get the Quote of the Day. Please try again."))
				return
			}
			quote := "Redisden gelen random string: " + quoteResp.Contents.Quotes[0].Quote
			client.Set(date, quote, 24*time.Hour)
			w.Write([]byte(quote))
		} else {
			log.Println("Cache Hit for date ", date)
			w.Write([]byte(val))
		}
	}
}
