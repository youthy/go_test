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
)

// connect db and init maxID
func init() {
	_, err := readConfig()
	fmt.Println(GlobalConfig)
	checkErr(err)
	ConnectPG()
	maxID = getDbMaxID()
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", UserIndexHandler).Methods("GET")
	router.HandleFunc("/users", UserCreateHandler).Methods("POST")
	router.HandleFunc("/users/{user_id}/relationships", RelationIndexHandler).
		Methods("Get")
	router.HandleFunc("/users/{user_id}/relationships/{other_id}", RelationUpdateHandler).
		Methods("PUT")
	err := http.ListenAndServe(fmt.Sprintf(":%v", GlobalConfig.ListenPort), router)
	checkErr(err)
}
