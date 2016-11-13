package handlers

import (
	"encoding/json"
	"jimmify-server/db"
	"net/http"
)

//Recent: get recent answered queries
func Recent(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})

	//Get recent
	recents, err := db.GetRecent(10)
	if err != nil {
		ReturnInternalServerError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response["recents"] = recents
	response["status"] = "true"
	json.NewEncoder(w).Encode(response)
}
