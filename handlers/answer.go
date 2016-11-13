package handlers

import (
	"encoding/json"
	"jimmify-server/db"
	"net/http"
)

//Index: return a status true to tell if the app is live
func Answer(w http.ResponseWriter, r *http.Request) {
	var q db.Query
	response := make(map[string]interface{})

	//read json
	err := json.NewDecoder(r.Body).Decode(q)
	if err != nil {
		ReturnStatusBadRequest(w, "Failed to decode query json")
		return
	}

	//add query
	err = db.AnswerQuery(q.Key, q.Answer)
	if err != nil {
		ReturnInternalServerError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response["status"] = "true"
	json.NewEncoder(w).Encode(response)
}
