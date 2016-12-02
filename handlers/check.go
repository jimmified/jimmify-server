package handlers

import (
	"encoding/json"
	"jimmify-server/db"
	"net/http"
)

//Check : check if post is answered
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
		//return status false
		w.WriteHeader(http.StatusOK)
		response["status"] = "false"
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response["answer"] = a.Answer
	response["status"] = "true"
	json.NewEncoder(w).Encode(response)
}
