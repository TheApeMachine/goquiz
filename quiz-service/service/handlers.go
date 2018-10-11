package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/theapemachine/goquiz/quiz-service/dbclient"
)

// DBClient : Setup an instance of our embedded key/value store.
var DBClient dbclient.IBoltClient

// GetQuizzes : Response implementation for getting a list of quizzes.
func GetQuizzes(w http.ResponseWriter, r *http.Request) {
	quizzes, err := DBClient.QueryQuizzes()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, _ := json.Marshal(quizzes)

	writeJSONResponse(w, http.StatusOK, data)
}

// GetQuiz : Response implementation for getting a quiz.
func GetQuiz(w http.ResponseWriter, r *http.Request) {
	var quizID = mux.Vars(r)["quizID"]

	quiz, err := DBClient.QueryQuiz(quizID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, _ := json.Marshal(quiz)

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
