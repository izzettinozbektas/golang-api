package handlers

import (
	"github.com/go-redis/redis"
	"github.com/izzettinozbektas/golang-api/internal/helpers"
	"github.com/izzettinozbektas/golang-api/internal/response"
	"log"
	"net/http"
	"time"
)

// redis client connect
var client = helpers.ConnetToRedis()

// redis the handler
func Redis(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)

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
		resp["data"] = quote
	} else {
		log.Println("Cache Hit for date ", date)
		resp["data"] = val
	}
	response.Write(w, response.Success("", resp), response.Code(http.StatusOK))
}
