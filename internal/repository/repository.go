package repository

import (
	"context"
	"restful-medods/internal/models"
)

type Repository interface {
	CreatePatient(ctx context.Context, patient models.Patient) (int64, error)
	CreateConsultationRequest(ctx context.Context, request models.ConsultationRequest) (int64, error)
	CreateRecommendation(ctx context.Context, recommendation models.Recommendation) (int64, error)
	GetRecommendationList(ctx context.Context, patientId int64) ([]models.PatientRecommendation, error)
}

var repository Repository

func NewRepository(repo Repository) {
	repository = repo
}

func CreatePatient(ctx context.Context, patient models.Patient) (int64, error) {
	return repository.CreatePatient(ctx, patient)
}
func CreateConsultationRequest(ctx context.Context, request models.ConsultationRequest) (int64, error) {
	return repository.CreateConsultationRequest(ctx, request)
}
func CreateRecommendation(ctx context.Context, recommendation models.Recommendation) (int64, error) {
	return repository.CreateRecommendation(ctx, recommendation)
}
func GetRecommendationList(ctx context.Context, patientId int64) ([]models.PatientRecommendation, error) {
	return repository.GetRecommendationList(ctx, patientId)
}
