package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"
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

// UserIndexHandler list user TODO divided into pages?
func UserIndexHandler(w http.ResponseWriter, r *http.Request) {
	defer serverErr(w)
	userList := make([]User, 0)
	err := userdb.Model(&userList).Select()
	fmt.Println("user", userList)
	checkErr(err)
	json.NewEncoder(w).Encode(&userList)
}

// UserCreateHandler create user
func UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	defer serverErr(w)
	// TODO when max ID
	if maxID >= MaxUserID {
		badRequest(w)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var newUser User
	err := decoder.Decode(&newUser)
	checkErr(err)
	// now default type is user. maybe admin and vip etc.. afterwards
	newUser.Type = "user"
	newUser.Timestamp = uint32(time.Now().Unix())
	newUser.ID = atomic.AddUint32(&maxID, 1)
	fmt.Println("user add", newUser)
	err = userdb.Insert(&newUser)
	checkErr(err)
	json.NewEncoder(w).Encode(&newUser)
}

func checkUserExist(id ID) error {
	user := new(User)
	err := userdb.Model(user).Where("id = ?", id).Select()
	return err
}
