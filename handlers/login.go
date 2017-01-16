package handlers

import (
	"encoding/json"
	"jimmify-server/auth"
	"net/http"
)

//Password type
type Credentials struct {
	Username string
	Password string
}

//Login login a user
func Login(w http.ResponseWriter, r *http.Request) {
	var c Credentials
	response := make(map[string]interface{})

	//read json
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		ReturnStatusBadRequest(w, "Failed to decode query json")
		return
	}

	//validate data
	err = validateLogin(c)
	if err != nil {
		ReturnUnauthorized(w, err.Error())
		return
	}

	//create token
	token, err := auth.CreateToken(c.Username)
	if err != nil {
		ReturnInternalServerError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response["token"] = token
	response["status"] = "true"
	json.NewEncoder(w).Encode(response)
}
