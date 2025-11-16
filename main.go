package main

import (
	"io"
	"log"
	"os"
	"study/service/config"
	app "study/service/svcHandler"

	"github.com/kardianos/service"
)

func main() {

	if err := config.LoadConfig(); err != nil {
		log.Fatal(err)
	}

	logFile, err := os.OpenFile(config.AppConfig.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	mw := io.MultiWriter(os.Stderr, logFile)

	log.SetOutput(mw)
	if err != nil {
		log.Fatalf("Failed to open log file %v", err)
	}
	defer logFile.Close()

	if config.AppConfig.ServiceName == "" {
		log.Println("No name set, configured to default Hummus")
		config.AppConfig.ServiceName = "Hummus"
	}

	prg := app.NewProgram(config.AppConfig)

	svcConfig := &service.Config{
		Name:        config.AppConfig.ServiceName,
		DisplayName: config.AppConfig.ServiceName,
		Description: "Very simple sauce to add some flavor and protection",
	}

	s, err := service.New(prg, svcConfig)

	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) > 1 {
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Fatalf("Error from control: %s", err)
		}
		return
	}

	err = s.Run()

	if err != nil {
		log.Fatalf("Error from service: %s", err)
	}
}
