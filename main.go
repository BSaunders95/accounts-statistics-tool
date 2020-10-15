package main

import (
	"fmt"
	"github.com/BSaunders95/accounts-statistics-tool/config"
	"github.com/BSaunders95/accounts-statistics-tool/service"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {

	cfg, err := config.Get()
	if err != nil {
		log.Error(fmt.Sprintf("Error when establishing config: %s", err))
		os.Exit(1)
	}

	svc := service.NewService(cfg)

	svc.GetNumberOfCICReports()
}
