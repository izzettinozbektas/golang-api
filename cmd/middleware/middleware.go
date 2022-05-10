package middleware

import (
	"github.com/izzettinozbektas/golang-api/internal/driver"
	"github.com/izzettinozbektas/golang-api/internal/handlers"
	"github.com/izzettinozbektas/golang-api/internal/response"
	"net/http"
	"strings"
	"time"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//gecici kullanım
		dbName := "golang-db"
		dbUname := "golang"
		dbPass := "golangpass"
		dbHost := "app-mysql" // mysql container name olmalı

		connection, _ := driver.ConnectSQL(dbHost, dbUname, dbPass, dbName)
		//strings.ReplaceAll(authToken,"Bearer ","")
		authToken := r.Header.Get("Authorization")
		tokenControl, _ := handlers.ConnectionToDB(connection).DB.TokenControl(strings.ReplaceAll(authToken, "Bearer ", ""))

		if tokenControl.ExpDate.Format("2006-01-02 15:04:05") < time.Now().Format("2006-01-02 15:04:05") {
			response.Write(w, response.Error("UnAuthorized", nil), response.Code(http.StatusUnauthorized))
			return
		}
		next.ServeHTTP(w, r)

	})
}
