package transport

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"restful-medods/internal/models"
	"restful-medods/internal/service"
	"restful-medods/pkg/util"
	"strconv"

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
	ctx := r.Context()

	var request struct {
		PatientFirstName   string `json:"first_name"`
		PatientMiddleName  string `json:"middle_name"`
		PatientLastName    string `json:"last_name"`
		PatientDateOfBirst string `json:"DateOfBirst"`
		PatientPhone       string `json:"phone"`
		PatientEmail       string `json:"email"`
		Text               string `json:"text"`
	}

	type response struct {
		ID      int64  `json:"id"`
		Message string `json:"message"`
	}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&request); err != nil {
		util.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	log.Printf("Req %+v", request)

	id, err := h.service.CreatePatient(ctx, models.Patient{
		FirstName:   request.PatientFirstName,
		MiddleName:  request.PatientMiddleName,
		LastName:    request.PatientLastName,
		Email:       request.PatientEmail,
		DateOfBirst: request.PatientDateOfBirst,
		Phone:       request.PatientPhone,
	})
	if err != nil {
		util.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	id, err = h.service.CreateConsultationRequest(ctx, models.ConsultationRequest{
		PatientID: id,
		Text:      request.Text,
	})
	if err != nil {
		util.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	util.ResponseOk(w, response{
		ID:      id,
		Message: "success",
	})
}

func (h *handler) CreateRecommendation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	type response struct {
		ID      int64  `json:"id"`
		Message string `json:"message"`
	}

	var request struct {
		Text string `json:"text"`
	}

	idStr := mux.Vars(r)["id"]
	if idStr == "" {
		util.ResponseError(w, http.StatusBadRequest, "id is empty")
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		util.ResponseError(w, http.StatusBadRequest, "id is not number")
		return
	}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&request)
	if err != nil {
		util.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	id, err = h.service.CreateRecommendation(ctx, models.Recommendation{
		ConsultationRequestID: id,
		Text:                  request.Text,
	})
	if err != nil {
		util.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	util.ResponseOk(w, response{
		ID:      id,
		Message: "success",
	})
}

func (h *handler) GetPatientRecommendations(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	type response struct {
		Recommendations []models.PatientRecommendation `json:"recommendations"`
		Total           int                            `json:"total"`
	}

	idStr := mux.Vars(r)["id"]
	if idStr == "" {
		util.ResponseError(w, http.StatusBadRequest, "id is empty")
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 0)
	if err != nil {
		util.ResponseError(w, http.StatusBadRequest, "id is not number")
		return
	}
	recommendations, err := h.service.GetPatientRecommendations(ctx, id)
	if err != nil {
		util.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := http.Get("https://api.fda.gov/drug/ndc.json?limit=10")
	if err != nil {
		log.Printf("http.Get medicine err %v", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("io.ReadAll err %v", err)
	}

	var data models.MedicineResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("json.Unmarshal err %v", err)
	}

	for i := range recommendations {
		recommendations[i].Medicines = data.Results[:rand.Intn(10)]
	}

	util.ResponseOk(w, response{
		Recommendations: recommendations,
		Total:           len(recommendations),
	})
}
