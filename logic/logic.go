package logic

import (
	"errors"

	"github.com/briankliwon/simple-golang-mock-unit-test/database"
	"github.com/briankliwon/simple-golang-mock-unit-test/model"
)

// Servie parent struct to handle all processing request
type UserService struct {
	Storage database.UserStorage
}

// Get service: goal => get user data by name user
func (user *UserService) Get(name string) (*model.User, error) {
	result := user.Storage.GetByName(name)
	if result == nil {
		return nil, errors.New("data is undefined")
	} else {
		return result, nil
	}
}
