package postgres

import (
	"context"
	"database/sql"
	"log"
	"restful-medods/internal/models"
	"restful-medods/pkg/client/postgresql"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(url string) (*Postgres, error) {
	db, err := postgresql.NewClient(url)
	if err != nil {
		log.Printf("postgresql.NewClient err %v", err)
		return nil, err
	}

	return &Postgres{
		db: db,
	}, nil
}

func (p *Postgres) CreatePatient(ctx context.Context, patient models.Patient) error {
	return nil
}

func (p *Postgres) CreateConsultationRequest(ctx context.Context, request models.ConsultationRequest) error {
	return nil
}

func (p *Postgres) CreateRecommendation(ctx context.Context, recommendation models.Recommendation) error {
	return nil
}

func (p *Postgres) GetRecommendationList(ctx context.Context, patientId int64) error {
	return nil
}
