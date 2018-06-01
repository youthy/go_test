package models

import (
	"fmt"

	"github.com/go-pg/pg"
	"go_test/config"
)

// Userdb is pointer to pg.DB
var Userdb *pg.DB

// ConnectPG use for connect postgresql
func ConnectPG() {
	if Userdb == nil {
		Userdb = pg.Connect(&pg.Options{
			Addr:     config.GlobalConfig.Host + ":" + fmt.Sprint(config.GlobalConfig.DbPort),
			User:     config.GlobalConfig.User,
			Password: config.GlobalConfig.Password,
			Database: config.GlobalConfig.Database,
		})
	}
}
