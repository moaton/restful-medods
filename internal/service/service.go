package service

import "restful-medods/internal/repository"

type Service interface {
	CreateConsultationRequest() error
	CreateRecommendation() error
	GetPatientRecommendations() error
}

type service struct {
	db repository.Repository
}

func NewService(db repository.Repository) Service {
	return &service{
		db: db,
	}
}

func (s *service) CreateConsultationRequest() error {
	return nil
}
func (s *service) CreateRecommendation() error {
	return nil
}
func (s *service) GetPatientRecommendations() error {
	return nil
}
