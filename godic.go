package main

import (
	"os"
	"path"

	log "github.com/Sirupsen/logrus"
	"github.com/go-martini/martini"
	"github.com/ikeikeikeike/godic/views"
)

func init() {
	initLogger()
}

func initLogger() {
	if martini.Env == "production" {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.WarnLevel)

		p, _ := os.Getwd()
		file, err := os.OpenFile(
			path.Join(p, "logs/godic.log"),
			os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666,
		)
		if err != nil {
			panic(err)
		}
		log.SetOutput(file)
	} else {
		log.SetFormatter(&log.TextFormatter{})
		log.SetLevel(log.DebugLevel)
		log.SetOutput(os.Stderr)
	}

	views.App.Map(log.StandardLogger())
}

func main() {
	views.App.Run()
}
