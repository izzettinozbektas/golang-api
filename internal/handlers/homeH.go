package handlers

import (
	"github.com/izzettinozbektas/golang-api/internal/response"
	"net/http"
)

// home is the handler
func Home(w http.ResponseWriter, r *http.Request) {
	response.Write(w, response.Success("Welcome! Please hit the `/redis` API to get the quote of the day.", nil), response.Code(http.StatusOK))
}
