package main

import (
	"errors"
	"github.com/burakkuru5534/src/api"
	"github.com/burakkuru5534/src/auth"
	"github.com/burakkuru5534/src/helper"
	"github.com/burakkuru5534/src/service"
	_ "github.com/lib/pq"
)

func main() {

	helper.InitConf()

	conInfo := helper.PgConnectionInfo{
		Host:     "127.0.0.1",
		Port:     5432,
		Database: "soft-robotics",
		Username: "postgres",
		Password: "tayitkan",
		SSLMode:  "disable",
	}

	helper.Conf.Auth = auth.NewAuth("2GcQCe7SuKxbaA3NSMBy8ztBPbfDsXJ4", false)

	db, err := helper.NewPgSqlxDbHandle(conInfo, 10)
	if err != nil {
		errors.New("create db handle error.")
	}
	err = db.Ping()
	if err != nil {
		errors.New("ping db error.")
	}

	// Create Appplication Service
	err = helper.InitApp(db)
	if err != nil {
		errors.New("init app error.")
	}

	service.StartHttpService(8080, api.HttpService())
}
