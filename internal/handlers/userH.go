package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/izzettinozbektas/golang-api/internal/helpers"
	"github.com/izzettinozbektas/golang-api/internal/models"
	"log"
	"net/http"
	"strconv"
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

	resp := make(map[string]string)
	if err != nil {
		log.Fatal(err)
	}
	resp["message"] = "Kayıt Başarılı"

	jresp, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jresp)
}
func (m *Repository) Users(w http.ResponseWriter, r *http.Request) {
	users, err := m.DB.Users()
	if err != nil {
		log.Fatal(err)
	}
	resp := make(map[string]interface{})
	resp["users"] = users

	jresp, jerr := json.Marshal(resp)
	if jerr != nil {
		log.Fatal(jerr)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jresp)

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
	resp := make(map[string]string)
	if err != nil {
		resp["message"] = err.Error()
	} else {
		resp["message"] = "İşlem Başarılı"
	}

	jresp, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jresp)
}
func (m *Repository) User(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	user, err := m.DB.User(id)
	resp := make(map[string]interface{})
	if err != nil {
		resp["message"] = err.Error()
	} else {
		resp["user"] = user
	}

	jresp, jerr := json.Marshal(resp)
	if jerr != nil {
		log.Fatal(jerr)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jresp)

}
func (m *Repository) UserDelete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	status, err := m.DB.UserDelete(id)
	if err != nil {
		log.Fatal(err)
	}
	resp := make(map[string]interface{})

	if status == true {
		resp["message"] = "işlem başarılı"
	}

	jresp, jerr := json.Marshal(resp)
	if jerr != nil {
		log.Fatal(jerr)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jresp)

}
