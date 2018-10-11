package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/theapemachine/goquiz/account-service/dbclient"
)

// DBClient : Setup an instance of our embedded key/value store.
var DBClient dbclient.IBoltClient

// GetAccount : Response implementation for getting an account.
func GetAccount(w http.ResponseWriter, r *http.Request) {
	var accountID = mux.Vars(r)["accountID"]

	account, err := DBClient.QueryAccount(accountID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, _ := json.Marshal(account)
	writeJSONResponse(w, http.StatusOK, data)
}

// HealthCheck : Endpoint handler to see the status of the service.
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	dbUp := DBClient.Check()

	if dbUp {
		data, _ := json.Marshal(healthCheckResponse{Status: "UP"})
		writeJSONResponse(w, http.StatusOK, data)
	} else {
		data, _ := json.Marshal(healthCheckResponse{Status: "Database not accessible"})
		writeJSONResponse(w, http.StatusServiceUnavailable, data)
	}
}

// writeJSONResponse : Generic method to write data as a JSON response, including the needed headers.
func writeJSONResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(status)
	w.Write(data)
}

type healthCheckResponse struct {
	Status string `json:"status"`
}
