package handlers

import (
	"encoding/json"
	"net/http"
)

//Index: return a status true to tell if the app is live
func Index(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	w.WriteHeader(http.StatusOK)
	response["status"] = "true"
	json.NewEncoder(w).Encode(response)
}
