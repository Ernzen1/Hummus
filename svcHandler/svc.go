package svcHandler

import (
	"log"
	"study/service/backupEngine"
	bconfig "study/service/config"
	"time"

	"github.com/kardianos/service"
)

type program struct {
	quit chan struct{}
}

func (p *program) Start(s service.Service) error {
	log.Println("Starting service")

	p.quit = make(chan struct{})

	go p.runBackupLoop()
	return nil

}

func (p *program) runBackupLoop() {
	log.Println("Starting backup loop, running a backup every %d seconds", bconfig.AppConfig.Tick)

	ticker := time.NewTicker((time.Duration(bconfig.AppConfig.Tick) * time.Second))

	defer ticker.Stop()

	log.Println("Starting backup for the first cycle")
	p.RunBack
}
