package database

import "github.com/briankliwon/simple-golang-mock-unit-test/model"

// interface UserStorage it can be used in Mock & real transaction
type UserStorage interface {
	GetByName(name string) *model.User
}
