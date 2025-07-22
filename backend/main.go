package main

import (
	"os"
	"os/signal"
	"sub/api"
	"sub/db"
	"sub/submitter"
	"sub/utils/config"
	"sub/utils/log"
)

func main() {
	config, err := config.LoadConfig("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	log.SetLogLevel(config.LogLevel)
	log.Info("Loaded config")
	log.Debugf("Config: %+v", config)

	db.InitDB(config.Database)
	defer db.CloseDB()
	db.ExecSQLFile("db/schema.sql")
	db.InitStatements("db/statements.sql")

	go submitter.Loop(config)
	go api.ServeAPI(config)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Info("Submitter Running. Press CTRL-C to exit")
	<-stop
	log.Info("Stopping submitter")
}
