/*package main serve as a RESTful HTTP server which provide functions such as
UserIndexHandler return list of users registered
UserCreateHandler can create user
RelationIndexHandler return one user's all relationships
RelationUpdateHandler update relations between user and another user
*/
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go_test/config"
	"go_test/controllers"
	"go_test/models"
	"go_test/util"
)

// Init connect db and init maxID
func Init() {
	_, err := config.ReadConfig()
	fmt.Println(config.GlobalConfig)
	util.CheckErr(err)
	models.ConnectPG()
}

func main() {
	Init()
	router := mux.NewRouter()
	router.HandleFunc("/users", controllers.UserIndexHandler).Methods("GET")
	router.HandleFunc("/users", controllers.UserCreateHandler).Methods("POST")
	router.HandleFunc("/users/{user_id}/relationships",
		controllers.RelationIndexHandler).Methods("Get")
	router.HandleFunc("/users/{user_id}/relationships/{other_id}",
		controllers.RelationUpdateHandler).Methods("PUT")
	err := http.ListenAndServe(fmt.Sprintf(":%v", config.GlobalConfig.ListenPort),
		router)
	util.CheckErr(err)
}
