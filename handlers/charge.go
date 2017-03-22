package handlers

import (
	"encoding/json"
	"jimmify-server/stripe-wrapper"
	"net/http"
)

//Charge move the person up the queue when they pay
func Charge(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})

	token := r.FormValue("stripeToken")
	query := r.FormValue("queryId")

	//add charge
	err := stripe.PrioritizeQuestion(token, query)

	if err != nil {
		ReturnInternalServerError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response["status"] = "true"
	json.NewEncoder(w).Encode(response)
}
