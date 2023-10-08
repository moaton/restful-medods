package repository

import (
	"context"
	"restful-medods/internal/models"
)

type Repository interface {
	CreatePatient(ctx context.Context, patient models.Patient) error
	CreateConsultationRequest(ctx context.Context, request models.ConsultationRequest) error
	CreateRecommendation(ctx context.Context, recommendation models.Recommendation) error
	GetRecommendationList(ctx context.Context, patientId int64) error
}

var repository Repository

func NewRepository(repo Repository) {
	repository = repo
}

func CreatePatient(ctx context.Context, patient models.Patient) error {
	return repository.CreatePatient(ctx, patient)
}
func CreateConsultationRequest(ctx context.Context, request models.ConsultationRequest) error {
	return repository.CreateConsultationRequest(ctx, request)
}
func CreateRecommendation(ctx context.Context, recommendation models.Recommendation) error {
	return repository.CreateRecommendation(ctx, recommendation)
}
func GetRecommendationList(ctx context.Context, patientId int64) error {
	return repository.GetRecommendationList(ctx, patientId)
}
