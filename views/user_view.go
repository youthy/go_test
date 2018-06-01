package views

import (
	"go_test/models"
)

// UserIndex return slice of users
func UserIndex(users []models.User) []models.User {
	for i := range users {
		users[i] = UserShow(users[i])
	}
	return users
}

// UserShow is view of user
func UserShow(user models.User) models.User {
	// do nothing
	return user
}
