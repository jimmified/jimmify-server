package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

//ReturnInternalServerError returns an Internal Server Error
func ReturnInternalServerError(w http.ResponseWriter, message string) {
	log.Println(message)
	response := make(map[string]string)
	response["Status"] = "false"
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(response)
}

//ReturnStatusBadRequest returns a Bad Request Error
func ReturnStatusBadRequest(w http.ResponseWriter, message string) {
	log.Println(message)
	response := make(map[string]string)
	response["Status"] = "false"
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(response)
}
