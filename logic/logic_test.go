package logic

import (
	"testing"

	"github.com/briankliwon/simple-golang-mock-unit-test/database"
	"github.com/briankliwon/simple-golang-mock-unit-test/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userStorage = &database.UserStorageMock{Mock: mock.Mock{}}
var userService = UserService{Storage: userStorage}

// Test when get return is not found
func TestUserGet_NotFound(t *testing.T) {
	userStorage.Mock.On("GetByName", "1").Return(nil)
	user, _ := userService.Get("1")
	assert.Nil(t, user)
}

// Test when get return is found
func TestUserGet_Found(t *testing.T) {
	//Dummy return user data
	user := model.User{
		Id:   1,
		Name: "Brian",
	}

	//Set to mocking struct to give dummy result when passing Brian id
	userStorage.Mock.On("GetByName", "Brian").Return(user)
	//Get data form service function
	userResullt, err := userService.Get("Brian")
	//Checking error on get service
	assert.Nil(t, err)
	//Checking reuslt of user is equal or not
	assert.Equal(t, &user, userResullt, "User is not same")
}
