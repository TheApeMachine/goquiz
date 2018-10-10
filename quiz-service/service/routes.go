package service

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetQuiz",
		"Get",
		"/quizzes/{quizID}",
		GetQuiz,
	},

	Route{
		"HealthCheck",
		"GET",
		"/health",
		HealthCheck,
	},
}
