package handlers

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/izzettinozbektas/golang-api/internal/helpers"
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
		resp["message"] = quote
	} else {
		log.Println("Cache Hit for date ", date)
		resp["message"] = val
	}

	jresp, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jresp)
}
