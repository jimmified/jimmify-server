package handlers

import (
	"encoding/json"
	"jimmify-server/stripe-wrapper"
	"net/http"
)

type ChargeType struct {
	Key   int64  `json:"key"`
	Token string `json:"token"`
}

//Charge move the person up the queue when they pay
func Charge(w http.ResponseWriter, r *http.Request) {
	var c ChargeType
	response := make(map[string]interface{})

	//read json
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		ReturnStatusBadRequest(w, "Failed to decode query json")
		return
	}

	//add charge
	err = stripe.PrioritizeQuestion(c.Token, c.Key)

	if err != nil {
		ReturnInternalServerError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response["status"] = true
	json.NewEncoder(w).Encode(response)
}
