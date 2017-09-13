package handlers

import (
    "encoding/json"
    "jimmify-server/db"
    "jimmify-server/auth"
    "net/http"
)

type ExpoMsg struct {
    UUID string    `json:"UUID"`
    Token string   `json:"token"`
    ExpoID string  `json:"expoID"`
}

func ExpoRegister(w http.ResponseWriter, r *http.Request) {
    var msg ExpoMsg
    response := make(map[string]interface{})

    err := json.NewDecoder(r.Body).Decode(&msg)
    if err != nil {
        ReturnStatusBadRequest(w, "Failed to decode json")
        return
    }

    _, err = auth.CheckToken(msg.Token)
    if err != nil {
        ReturnUnauthorized(w, err.Error())
        return
    }

    err = db.AddExpoClient(msg.UUID, msg.ExpoID)

    if err != nil {
        ReturnInternalServerError(w, err.Error())
        return
    }

    w.WriteHeader(http.StatusOK)
    response["status"] = true
    json.NewEncoder(w).Encode(response)
}

func ExpoUnRegister(w http.ResponseWriter, r *http.Request) {
    var msg ExpoMsg
    response := make(map[string]interface{})

    err := json.NewDecoder(r.Body).Decode(&msg)
    if err != nil {
        ReturnStatusBadRequest(w, "Failed to decode json")
        return
    }

    _, err = auth.CheckToken(msg.Token)
    if err != nil {
        ReturnUnauthorized(w, err.Error())
        return
    }

    err = db.RemoveExpoClient(msg.UUID)

    if err != nil {
        ReturnInternalServerError(w, err.Error())
        return
    }

    w.WriteHeader(http.StatusOK)
    response["status"] = true
    json.NewEncoder(w).Encode(response)
}
