package handlers

import (
	"encoding/json"
	"jimmify-server/db"
	"jimmify-server/firebase"
	"net/http"
)

//PushEnabled whether or not to send push
var PushEnabled bool

//Query : submit a query
func Query(w http.ResponseWriter, r *http.Request) {
	var q db.Query
	response := make(map[string]interface{})

	//read json
	err := json.NewDecoder(r.Body).Decode(&q)
	if err != nil {
		ReturnStatusBadRequest(w, "Failed to decode query json")
		return
	}

	//validate data
	err = validateQuery(q)
	if err != nil {
		ReturnStatusBadRequest(w, err.Error())
		return
	}

	//add query
	isDupe, d := db.IsDuplicate(q.Text)
	key := d.Key
	if !isDupe {
		key, err = db.AddQuery(q)
		if err != nil {
			ReturnInternalServerError(w, err.Error())
			return
		}
	}

	if PushEnabled == true {
		firebase.Push("Jimmy Query", q.Text)
	}

	w.WriteHeader(http.StatusOK)
	response["key"] = key
	response["status"] = true
	json.NewEncoder(w).Encode(response)
}
