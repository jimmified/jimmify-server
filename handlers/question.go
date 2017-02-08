package handlers

import (
    "encoding/json"
    "jimmify-server/db"
    "net/http"
    "log"
)

//Question : get question by query id
func Question(w http.ResponseWriter, r *http.Request) {
    var q db.Query
    response :=make(map[string]interface{})

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
    log.Println(q.Key)
    //get question
    a, err := db.GetQuestion(q.Key)
    log.Println(a)
    log.Println(err)
    if err != nil {
        //return status false
        w.WriteHeader(http.StatusOK)
        response["status"] = "false"
        json.NewEncoder(w).Encode(response)
        return
    }

    w.WriteHeader(http.StatusOK)
    response["key"] = a.Key
    response["text"] = a.Text
    response["type"] = a.Type
    response["status"] = "true"
    json.NewEncoder(w).Encode(response)
}
