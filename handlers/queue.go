package handlers

import (
	"encoding/json"
	"jimmify-server/db"
	"net/http"
)

//Index: return a status true to tell if the app is live
func Queue(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})

	//add query
	queries, err := db.GetQueue(10)
	if err != nil {
		ReturnInternalServerError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response["queue"] = queries
	response["status"] = "true"
	json.NewEncoder(w).Encode(response)
}
