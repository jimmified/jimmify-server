package handlers

import (
	"encoding/json"
	"jimmify-server/auth"
	"jimmify-server/db"
	"net/http"
)

//Answer: let jimmy answer queries
func Answer(w http.ResponseWriter, r *http.Request) {
	var q db.Query
	response := make(map[string]interface{})

	//read json
	err := json.NewDecoder(r.Body).Decode(&q)
	if err != nil {
		ReturnStatusBadRequest(w, "Failed to decode query json")
		return
	}

	//check token
	_, err = auth.CheckToken(q.Token)
	if err != nil {
		ReturnUnauthorized(w, err.Error())
		return
	}

	//validate data
	err = validateAnswer(q)
	if err != nil {
		ReturnStatusBadRequest(w, err.Error())
		return
	}

	//append and convert list
	switch q.Type {
	case "search":
		q.List = append(q.List, db.RandomLinks()...)
	}
	linkStr, err := json.Marshal(q.List)

	//add query
	err = db.AnswerQuery(q.Key, q.Answer, string(linkStr))
	if err != nil {
		ReturnInternalServerError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response["status"] = true
	json.NewEncoder(w).Encode(response)
}
