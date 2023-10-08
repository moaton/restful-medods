package transport

import (
	"net/http"
	"restful-medods/internal/service"

	"github.com/gorilla/mux"
)

type Handler interface {
	CreateConsultationRequest(w http.ResponseWriter, r *http.Request)
	CreateRecommendation(w http.ResponseWriter, r *http.Request)
	GetPatientRecommendations(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) (router *mux.Router) {
	router = mux.NewRouter()
	handler := &handler{
		service: service,
	}

	router.HandleFunc("/consultation-requests", handler.CreateConsultationRequest).Methods("POST")
	router.HandleFunc("/consultation-requests/{id}/recommendation", handler.CreateRecommendation).Methods("POST")
	router.HandleFunc("/patient/{id}/recommendations", handler.GetPatientRecommendations).Methods("GET")

	return router
}

func (h *handler) CreateConsultationRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *handler) CreateRecommendation(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *handler) GetPatientRecommendations(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
