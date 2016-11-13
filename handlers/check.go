package handlers

import (
	"encoding/json"
	"jimmify-server/db"
	"net/http"
)

//Check: check if post is answered
func Check(w http.ResponseWriter, r *http.Request) {
	var q db.Query
	response := make(map[string]interface{})

	//read json
	err := json.NewDecoder(r.Body).Decode(&q)
	if err != nil {
		ReturnStatusBadRequest(w, "Failed to decode query json")
		return
	}

	//add query
	a, err := db.CheckQuery(q.Key)
	if err != nil {
		ReturnStatusBadRequest(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response["answer"] = a
	response["status"] = "true"
	json.NewEncoder(w).Encode(response)
}
