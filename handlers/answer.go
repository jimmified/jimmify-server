package handlers

import (
	"encoding/json"
	"jimmify-server/auth"
	"jimmify-server/db"
	"math/rand"
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

	// check if list was provided
	linkStr, err := json.Marshal(q.List)

	//add query
	err = db.AnswerQuery(q.Key, q.Answer, string(linkStr))
	if err != nil {
		ReturnInternalServerError(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response["status"] = "true"
	json.NewEncoder(w).Encode(response)
}

func RandomLink() string {
	var links [3]string
	links[0] = "https://www.youtube.com/watch?v=VxTQKxyJyxw"
	links[1] = "https://media.giphy.com/media/B6sl8C4moPBGo/giphy.gif"
	links[2] = "http://i.imgur.com/zrSoDU9.jpg"
	return links[rand.Intn(len(links))]
}
