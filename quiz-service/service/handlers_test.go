package service

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/theapemachine/goquiz/quiz-service/dbclient"
	"github.com/theapemachine/goquiz/quiz-service/model"
)

func TestGetQuizWrongPath(t *testing.T) {
	Convey("Given a HTTP request for /invalid/123", t, func() {
		req := httptest.NewRequest("GET", "/invalid/123", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 404", func() {
				So(resp.Code, ShouldEqual, 404)
			})
		})
	})
}

func TestGetQuiz(t *testing.T) {
	// Get a mock instance of the IBoltClient interface.
	mockRepo := &dbclient.MockBoltClient{}

	mockRepo.On("QueryQuiz", "123").Return(model.Quiz{ID: "123", Name: "Person_123"}, nil)
	mockRepo.On("QueryQuiz", "456").Return(model.Quiz{}, fmt.Errorf("Some error"))

	DBClient = mockRepo

	Convey("Given a HTTP request for /quizzes/123", t, func() {
		req := httptest.NewRequest("GET", "/quizzes/123", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 200", func() {
				So(resp.Code, ShouldEqual, 200)

				quiz := model.Quiz{}
				json.Unmarshal(resp.Body.Bytes(), &quiz)
				So(quiz.ID, ShouldEqual, "123")
				So(quiz.Name, ShouldEqual, "Person_123")
			})
		})
	})

	Convey("Given a HTTP request for /quizzes/456", t, func() {
		req := httptest.NewRequest("GET", "/quizzes/456", nil)
		resp := httptest.NewRecorder()

		Convey("When the request is handled by the Router", func() {
			NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 404", func() {
				So(resp.Code, ShouldEqual, 404)
			})
		})
	})
}
