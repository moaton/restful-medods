package app

import (
	"fmt"
	"log"
	"net/http"
	"restful-medods/internal/models"
	"restful-medods/internal/repository"
	"restful-medods/internal/repository/postgres"
	"restful-medods/internal/service"
	"restful-medods/internal/transport"
)

func Run(cfg *models.Config) {
	fmt.Println("Run started")
	addr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresDB)
	repo, err := postgres.NewPostgres(addr)
	if err != nil {
		log.Fatalf("postgres.NewPostgres err %v", err)
	}
	repository.NewRepository(repo)

	service := service.NewService(repo)

	router := transport.NewHandler(service)

	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatalf("http.ListenAndServe err %v", err)
	}
}
