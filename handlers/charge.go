package handlers

import (
	"encoding/json"
	"jimmify-server/db"
	"jimmify-server/stripe"
	"net/http"
)

//Charge move the person up the queue when they pay
func Charge(w http.ResponseWriter, r *http.Request) {
	var c db.Charge
	response := make(map[string]interface{})

	//read json
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		ReturnStatusBadRequest(w, "Failed to decode charge json")
		return
	}

	//vaildate data
	err = validateCharge(c)
	if err != nil {
		ReturnStatusBadRequest(w, err.Error())
		return
	}

	//add charge
	err = stripe.PrioritizeQuestion(c.ID, c.Query)

	if err != nil {
		ReturnInternalServerError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response["status"] = "true"
	json.NewEncoder(w).Encode(response)
}
