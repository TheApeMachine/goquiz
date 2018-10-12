package dbclient

import (
	"github.com/stretchr/testify/mock"
	"github.com/theapemachine/goquiz/quiz-service/model"
)

// MockBoltClient : Setup a mock for the key/value store.
type MockBoltClient struct {
	mock.Mock
}

// QueryQuizzes : Implements a unit test for retrieving a list of quizzes.
// TODO: Implement the test.
func (m *MockBoltClient) QueryQuizzes() ([]model.Quiz, error) {
	return nil, nil
}

// QueryQuiz : Implements a unit test for retrieving a quiz by id.
func (m *MockBoltClient) QueryQuiz(quizID string) (model.Quiz, error) {
	args := m.Mock.Called(quizID)
	return args.Get(0).(model.Quiz), args.Error(1)
}

func (m *MockBoltClient) OpenBoltDb() {
	// Does nothing
}

func (m *MockBoltClient) Seed() {
	// Does nothing
}

func (m *MockBoltClient) Check() bool {
	args := m.Mock.Called()
	return args.Get(0).(bool)
}
