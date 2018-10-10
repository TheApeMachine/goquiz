package dbclient

import (
	"github.com/stretchr/testify/mock"
	"github.com/theapemachine/goquiz/account-service/model"
)

// MockBoltClient : Setup a mock for the key/value store.
type MockBoltClient struct {
	mock.Mock
}

func (m *MockBoltClient) QueryAccount(accountID string) (model.Account, error) {
	args := m.Mock.Called(accountID)
	return args.Get(0).(model.Account), args.Error(1)
}

func (m *MockBoltClient) OpenBoltDb() {
	// Does nothing
}

func (m *MockBoltClient) Seed() {
	// Does nothing
}
