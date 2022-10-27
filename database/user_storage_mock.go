package database

import (
	"github.com/briankliwon/simple-golang-mock-unit-test/model"
	"github.com/stretchr/testify/mock"
)

// parent mock variable to set & accessing all data test
type UserStorageMock struct {
	Mock mock.Mock
}

// replication of actual get user from database
func (storage *UserStorageMock) GetByName(name string) *model.User {
	args := storage.Mock.Called(name)
	if args.Get(0) == nil {
		return nil
	}
	result := args.Get(0).(model.User)
	return &result
}
