package postgres

import (
	"context"
	"log"
	"restful-medods/internal/models"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreatePatient(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("sqlmock.New err %v", err)
	}

	repository := Postgres{
		db: db,
	}

	patient := models.Patient{
		FirstName:   "Anet",
		MiddleName:  "",
		LastName:    "Zhuban",
		Email:       "anet@gj.com",
		Phone:       "+334442211",
		DateOfBirst: "12.04.1995",
	}
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery("^SELECT id FROM patients*").
		WithArgs(patient.FirstName, patient.LastName, patient.MiddleName).
		WillReturnRows(rows)

	rows = sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery("^INSERT INTO sms_fraud*").
		WithArgs(patient.FirstName, patient.LastName, patient.MiddleName, patient.Email, patient.Phone, patient.DateOfBirst).
		WillReturnRows(rows)

	id, err := repository.CreatePatient(context.Background(), patient)
	if err != nil {
		log.Fatalf("TestCreatePatient repository.CreatePatient err %v", err)
	}
	assert.Equal(t, int64(1), id, "they should be equal")
}

func TestCreateConsultationRequest(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("sqlmock.New err %v", err)
	}

	repository := Postgres{
		db: db,
	}

	createdAt := time.Now().UTC()
	request := models.ConsultationRequest{
		PatientID: 1,
		Text:      "test",
		CreatedAt: createdAt,
	}
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery("^INSERT INTO consultationrequest*").
		WithArgs(request.PatientID, request.Text, createdAt).
		WillReturnRows(rows)

	id, err := repository.CreateConsultationRequest(context.Background(), request)
	if err != nil {
		log.Fatalf("TestCreatePatient repository.CreatePatient err %v", err)
	}
	assert.Equal(t, int64(1), id, "they should be equal")
}

func TestCreateRecommendation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("sqlmock.New err %v", err)
	}

	repository := Postgres{
		db: db,
	}

	recommendation := models.Recommendation{
		ConsultationRequestID: 1,
		Text:                  "test",
	}
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery("^INSERT INTO recommendations*").
		WithArgs(recommendation.ConsultationRequestID, recommendation.Text).
		WillReturnRows(rows)

	id, err := repository.CreateRecommendation(context.Background(), recommendation)
	if err != nil {
		log.Fatalf("TestCreateRecommendation repository.CreateRecommendation err %v", err)
	}
	assert.Equal(t, int64(1), id, "they should be equal")
}

func TestGetRecommendationList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("sqlmock.New err %v", err)
	}

	repository := Postgres{
		db: db,
	}

	rows := sqlmock.NewRows([]string{"firstname", "middlename", "lastname", "text", "recommendations"}).AddRow("Anet", "", "Zhuban", "text", "recomendations")
	mock.ExpectQuery("^SELECT p.firstname, p.middlename, p.lastname, c.text as request_text, (.+) as recomendations FROM patients p  JOIN consultationrequest c (.+) JOIN recommendations r (.+)*").WithArgs(1).WillReturnRows(rows)

	patientRecomendations, err := repository.GetRecommendationList(context.Background(), 1)
	if err != nil {
		log.Fatalf("TestGetRecommendationList repository.GetRecommendationList err %v", err)
	}

	expected := []models.PatientRecommendation{
		{
			FirstName:       "Anet",
			MiddleName:      "",
			LastName:        "Zhuban",
			RequestText:     "text",
			Recommendations: "recomendations",
		},
	}
	assert.Equal(t, expected, patientRecomendations, "they should be equal")
}
