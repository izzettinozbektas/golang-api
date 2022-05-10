package handlers

import (
	"github.com/izzettinozbektas/golang-api/internal/helpers"
	"github.com/izzettinozbektas/golang-api/internal/models"
	"github.com/izzettinozbektas/golang-api/internal/response"
	"net/http"
)

func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {

	var authuser models.Authentication
	authuser.Email = r.FormValue("email")
	authuser.Password = r.FormValue("password")

	user, err := m.DB.Login(authuser)
	if err != nil {
		response.Write(w, response.Error("Email and Password is incorrect", nil), response.Code(http.StatusNotFound))
		return
	}
	if user.Email == "" {
		response.Write(w, response.Error("Email and Password is incorrect", nil), response.Code(http.StatusNotFound))
		return
	}
	check := helpers.CheckPasswordHash(authuser.Password, user.Password)

	if !check {
		response.Write(w, response.Error("Email and Password is incorrect", nil), response.Code(http.StatusNotFound))
		return
	}
	validToken, err := helpers.GenerateJWT(user.Email, string(user.AccessLevel))
	if err != nil {
		response.Write(w, response.Error("Token not created", nil), response.Code(http.StatusNotFound))
	}
	var token models.Token
	token.Role = user.AccessLevel
	token.Email = user.Email
	token.Token = validToken

	loginCreate, err := m.DB.LoginCreate(validToken)
	if err != nil {
		response.Write(w, response.Error("Login token not saved", map[string]string{"error": err.Error()}), response.Code(http.StatusNotFound))
		return
	}

	if loginCreate {
		response.Write(w, response.Success("Success", token), response.Code(http.StatusOK))
		return
	}

}
