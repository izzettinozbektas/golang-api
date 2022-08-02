package helpers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/izzettinozbektas/golang-api/internal/config"
	"github.com/izzettinozbektas/golang-api/internal/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func GetQuoteFromAPI() (*models.QuoteResponse, error) {
	API_URL := GetConfig().GetQuoteFromAPI
	resp, err := http.Get(API_URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	log.Println("Quote API Returned: ", resp.StatusCode, http.StatusText(resp.StatusCode))

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		quoteResp := &models.QuoteResponse{}
		json.NewDecoder(resp.Body).Decode(quoteResp)
		return quoteResp, nil
	} else {
		return nil, errors.New("Could not get quote from API")
	}

}
func GetConfig() config.Config {
	config, err := config.LoadConfig("build/")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	return config
}
func WaitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}

func ConnetToRedis() redis.Client {
	// Create Redis Client
	client := redis.NewClient(&redis.Options{
		Addr:     GetConfig().REDISURL,
		Password: GetConfig().REDISPASS,
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	return *client
}

func HashPassword(password string) string {
	hashpassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashpassword)
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func GenerateJWT(email, role string) (string, error) {
	var mySigningKey = []byte("can-you-keep-a-secret?")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).UTC().Format("2006-01-02 15:04:05")

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
func DecodeJWT(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecret := []byte("can-you-keep-a-secret?")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
