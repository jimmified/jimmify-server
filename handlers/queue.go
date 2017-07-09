package handlers

import (
	"encoding/json"
	"jimmify-server/db"
	"net/http"
)

//Queue: return top of queue
func Queue(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})

	//get queue
	queries, err := db.GetQueue(10)
	if err != nil {
		ReturnInternalServerError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response["queue"] = queries
	response["status"] = true
	json.NewEncoder(w).Encode(response)
}
