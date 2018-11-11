package main

import (
	"log"
	"logging/route"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	file, err := os.OpenFile("logger.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	defer file.Close()
	// logging := log.New(file, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)
	logging := logrus.New()
	logging.SetOutput(file)
	logging.SetFormatter(&logrus.JSONFormatter{})
	route := route.NewRoute(logging)
	route.Run(":3000")
}
