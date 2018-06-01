package models

import (
	"time"
)

// User struct
type User struct {
	//user      struct{} `sql:"alias:users"`
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Timestamp uint32 `json:"-"`
}

// ID alias uint32
type ID uint32

// CheckUserExist return err if a user_id is not in db
func CheckUserExist(id ID) error {
	user := new(User)
	err := Userdb.Model(user).Where("id = ?", id).Select()
	return err
}

// NewUser insert user into db. currently user type only have "user"
// mqybe "vip", "admin" afterwards.
func NewUser(user *User) error {
	user.Type = "user"
	user.Timestamp = uint32(time.Now().Unix())
	err := Userdb.Insert(user)
	return err
}

// GetUsers return all users in table users
func GetUsers() ([]User, error) {
	userList := make([]User, 0)
	err := Userdb.Model(&userList).Select()
	return userList, err
}
