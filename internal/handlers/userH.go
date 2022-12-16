package handlers

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/izzettinozbektas/golang-api/internal/helpers"
	"github.com/izzettinozbektas/golang-api/internal/models"
	"github.com/izzettinozbektas/golang-api/internal/response"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (m *Repository) UserCreate(w http.ResponseWriter, r *http.Request) {

	var user models.User
	user.FirstName = r.FormValue("first_name")
	user.LastName = r.FormValue("last_name")
	user.Email = r.FormValue("email")
	user.Password = helpers.HashPassword(r.FormValue("password"))
	user.AccessLevel = 1
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err := m.DB.UserCreate(user)
	if err != nil {
		log.Fatal(err)
	}
	response.Write(w, response.Success("işlem başarılı", nil), response.Code(http.StatusCreated))
}
func (m *Repository) Users(w http.ResponseWriter, r *http.Request) {
	users, err := m.DB.Users()
	if err != nil {
		log.Fatal(err)
	}
	resp := make(map[string]interface{})
	resp["users"] = users

	response.Write(w, response.Success("", resp), response.Code(http.StatusOK))

}
func (m *Repository) UserUpdate(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var user models.User

	user.FirstName = r.FormValue("first_name")
	user.LastName = r.FormValue("last_name")
	user.Email = r.FormValue("email")
	user.Password = helpers.HashPassword(r.FormValue("password"))
	user.UpdatedAt = time.Now()

	err := m.DB.UserUpdate(id, user)
	if err != nil {
		log.Fatal(err)
	}
	response.Write(w, response.Success("işlem başarılı", nil), response.Code(http.StatusOK))
}
func (m *Repository) User(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	user, err := m.DB.User(id)
	resp := make(map[string]interface{})
	if err != nil {
		response.Write(w, response.Error("işlem başarısız", map[string]string{"error": err.Error()}), response.Code(http.StatusNotFound))
	} else {
		resp["user"] = user
		response.Write(w, response.Success("", resp), response.Code(http.StatusOK))
	}
}
func (m *Repository) UserDelete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	status, err := m.DB.UserDelete(id)
	if err != nil {
		log.Fatal(err)
	}
	if status == true {
		response.Write(w, response.Success("işlem başarılı", nil), response.Code(http.StatusOK))
	}

}
func (m *Repository) GetUserInformation(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]interface{})

	authToken := r.Header.Get("Authorization")
	data, _ := helpers.DecodeJWT(strings.ReplaceAll(authToken, "Bearer ", ""))
	var token models.Token
	token.Email = fmt.Sprint(data["email"])
	user, err := m.DB.UserFromEmail(token.Email)

	if err != nil {
		response.Write(w, response.Error("işlem başarısız", map[string]string{"error": err.Error()}), response.Code(http.StatusNotFound))
	} else {
		resp["user"] = user
		response.Write(w, response.Success("", resp), response.Code(http.StatusOK))
	}
}
