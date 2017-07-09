package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

//ReturnInternalServerError returns an Internal Server Error
func ReturnInternalServerError(w http.ResponseWriter, message string) {
	log.Println(message)
	response := make(map[string]interface{})
	response["status"] = false
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(response)
}

//ReturnStatusBadRequest returns a Bad Request Error
func ReturnStatusBadRequest(w http.ResponseWriter, message string) {
	log.Println(message)
	response := make(map[string]interface{})
	response["status"] = false
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(response)
}

//ReturnUnauthorized returns a Unauthorized Error
func ReturnUnauthorized(w http.ResponseWriter, message string) {
	log.Println(message)
	response := make(map[string]interface{})
	response["Status"] = false
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(response)
}
