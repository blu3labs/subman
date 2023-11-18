package controller

import (
	"encoding/json"
	"net/http"
	"subman/database"
	"subman/types"
)

func PutSubscription(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body := r.Body
	defer body.Close()
	var subscription types.Subscription
	err := json.NewDecoder(body).Decode(&subscription)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	err = database.InsertSubscription(subscription)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetExecuteableSubscriptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	executeableSubscriptions, err := database.GetExecuteableSubscriptions()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(executeableSubscriptions)
}
