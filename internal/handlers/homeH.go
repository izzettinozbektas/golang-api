package handlers

import (
	"encoding/json"
	"net/http"
)

// home is the handler
func Home(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Welcome! Please hit the `/redis` API to get the quote of the day."

	jresp, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(jresp)
}
