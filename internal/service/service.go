package service

import (
	"context"
	"restful-medods/internal/models"
	"restful-medods/internal/repository"
)

type Service interface {
	CreatePatient(ctx context.Context, patient models.Patient) (int64, error)
	CreateConsultationRequest(ctx context.Context, req models.ConsultationRequest) (int64, error)
	CreateRecommendation(ctx context.Context, recommendation models.Recommendation) (int64, error)
	GetPatientRecommendations(ctx context.Context, id int64) ([]models.PatientRecommendation, error)
}

type service struct {
	db repository.Repository
}

func NewService(db repository.Repository) Service {
	return &service{
		db: db,
	}
}

func (s *service) CreatePatient(ctx context.Context, patient models.Patient) (int64, error) {
	id, err := s.db.CreatePatient(ctx, patient)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *service) CreateConsultationRequest(ctx context.Context, req models.ConsultationRequest) (int64, error) {
	id, err := s.db.CreateConsultationRequest(ctx, req)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *service) CreateRecommendation(ctx context.Context, recommendation models.Recommendation) (int64, error) {
	id, err := s.db.CreateRecommendation(ctx, recommendation)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *service) GetPatientRecommendations(ctx context.Context, id int64) ([]models.PatientRecommendation, error) {
	recomdendations, err := s.db.GetRecommendationList(ctx, id)
	if err != nil {
		return []models.PatientRecommendation{}, nil
	}
	return recomdendations, nil
}
