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

	//validate data
	err = validateCheck(q)
	if err != nil {
		ReturnStatusBadRequest(w, err.Error())
		return
	}

	//check query
	a, err := db.CheckQuery(q.Key)
	if err != nil {
		//return status false
		w.WriteHeader(http.StatusOK)
		response["status"] = "false"
		response["position"] = a.Position
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response["answer"] = a.Answer
	response["text"] = a.Text
	response["status"] = "true"
	json.NewEncoder(w).Encode(response)
}
