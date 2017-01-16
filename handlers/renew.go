package handlers

import (
	"encoding/json"
	"jimmify-server/auth"
	"jimmify-server/db"
	"net/http"
)

//Check : check if post is answered
func Renew(w http.ResponseWriter, r *http.Request) {
	var q db.Query
	response := make(map[string]interface{})

	//read json
	err := json.NewDecoder(r.Body).Decode(&q)
	if err != nil {
		ReturnStatusBadRequest(w, "Failed to decode query json")
		return
	}

	//check token
	user, err := auth.CheckToken(q.Token)
	if err != nil {
		ReturnUnauthorized(w, err.Error())
		return
	}

	//create token
	token, err := auth.CreateToken(user)
	if err != nil {
		ReturnInternalServerError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response["token"] = token
	response["status"] = "true"
	json.NewEncoder(w).Encode(response)
}
