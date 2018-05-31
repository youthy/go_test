package main

import (
	"fmt"

	"github.com/go-pg/pg"
)

var userdb *pg.DB
var maxID uint32

// ConnectPG use for connect postgresql
func ConnectPG() {
	if userdb == nil {
		userdb = pg.Connect(&pg.Options{
			Addr:     GlobalConfig.Host + ":" + fmt.Sprint(GlobalConfig.DbPort),
			User:     GlobalConfig.User,
			Password: GlobalConfig.Password,
			Database: GlobalConfig.Database,
		})
	}
}

func getDbMaxID() uint32 {
	var id uint32
	userdb.QueryOne(&id, `SELECT max(id) from users`)
	return id
}
