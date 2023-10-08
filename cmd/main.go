package main

import (
	"log"
	"os"
	"os/signal"
	"restful-medods/internal/app"
	"restful-medods/internal/models"
	"syscall"

	"github.com/kelseyhightower/envconfig"
)

func main() {
	var cfg models.Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("envconfig.Process err %v", err)
	}

	go app.Run(&cfg)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer signal.Stop(sigs)
	log.Println("Running...")
	<-sigs
	log.Println("Gracfull Shutdown")

}
