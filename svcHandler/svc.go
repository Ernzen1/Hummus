package svcHandler

import (
	"log"
	"study/service/backupEngine"
	bconfig "study/service/config"
	"time"

	"github.com/kardianos/service"
)

type Program struct {
	quit   chan struct{}
	config bconfig.Config
}

func NewProgram(cfg bconfig.Config) *Program {
	return &Program{
		config: cfg,
	}
}

func (p *Program) Start(s service.Service) error {
	log.Println("Starting service")

	p.quit = make(chan struct{})

	go p.runBackupLoop()
	return nil

}

func (p *Program) Stop(s service.Service) error {
	log.Println("Stopping service")
	close(p.quit)
	return nil
}

func (p *Program) runBackupLoop() {
	log.Printf("Starting backup loop, running a backup every %d seconds", bconfig.AppConfig.Tick)

	ticker := time.NewTicker((time.Duration(p.config.Tick) * time.Second))

	defer ticker.Stop()

	log.Println("Starting backup for the first cycle")
	p.runBackupLogic()

	for {
		select {
		case <-ticker.C:
			log.Println("New backup cycle detected!")
			p.runBackupLogic()

		case <-p.quit:
			log.Println("Quitting backup loop")
			return
		}

	}
}

func (p *Program) runBackupLogic() {
	backupEngine.BackupHandler(p.config.Paths, p.config.Location)
}
