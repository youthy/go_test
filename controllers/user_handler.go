package controllers

import (
	"encoding/json"
	"net/http"

	"go_test/models"
	"go_test/util"
	"go_test/views"
)

// UserIndexHandler list user TODO divided into pages?
func UserIndexHandler(w http.ResponseWriter, r *http.Request) {
	defer util.ServerErr(w)
	userList, err := models.GetUsers()
	util.CheckErr(err)
	views.Render(w, views.UserIndex(userList))
}

// UserCreateHandler create user
func UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	defer util.ServerErr(w)
	decoder := json.NewDecoder(r.Body)
	var newUser models.User
	err := decoder.Decode(&newUser)
	util.CheckErr(err)
	err = models.NewUser(&newUser)
	util.CheckErr(err)
	views.Render(w, views.UserShow(newUser))
}
