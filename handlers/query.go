package handlers

import (
	"encoding/json"
	"jimmify-server/db"
	"net/http"
)

//Index: return a status true to tell if the app is live
func Query(w http.ResponseWriter, r *http.Request) {
	var q db.Query
	response := make(map[string]interface{})
	response["status"] = "false"

	//read json
	err := json.NewDecoder(r.Body).Decode(q)
	if err != nil {
		ReturnStatusBadRequest(w, "Failed to decode query json")
		return
	}

	//add query
	key, err := db.AddQuery(q)
	if err != nil {
		ReturnStatusBadRequest(w, err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response["key"] = key
	response["status"] = "true"
	json.NewEncoder(w).Encode(response)
}
