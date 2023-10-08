package postgres

import (
	"context"
	"database/sql"
	"log"
	"restful-medods/internal/models"
	"restful-medods/pkg/client/postgresql"
	"time"

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

func (p *Postgres) CreatePatient(ctx context.Context, patient models.Patient) (int64, error) {
	var id int64
	row := p.db.QueryRowContext(ctx, "SELECT id FROM patients WHERE firstname = $1 AND lastname = $2 AND middlename = $3", patient.FirstName, patient.LastName, patient.MiddleName)
	err := row.Scan(&id)
	if err != nil {
		log.Println("CreatePatient get err ", err)
	}
	if id != 0 {
		return id, nil
	}

	err = p.db.QueryRowContext(ctx, "INSERT INTO patients (firstname, lastname, middlename, email, phone, dateofbirth) VALUES ($1, $2, $3, $4, $5, $6) RETURNING ID", patient.FirstName, patient.LastName, patient.MiddleName, patient.Email, patient.Phone, patient.DateOfBirst).Scan(&id)
	if err != nil {
		log.Println("CreatePatient insert err ", err)
		return 0, err
	}
	return id, nil
}

func (p *Postgres) CreateConsultationRequest(ctx context.Context, request models.ConsultationRequest) (int64, error) {
	var id int64
	err := p.db.QueryRowContext(ctx, "INSERT INTO consultationrequest (patientid, text, createdat) VALUES ($1, $2, $3) RETURNING ID", request.PatientID, request.Text, time.Now().UTC()).Scan(&id)
	if err != nil {
		log.Println("CreateConsultationRequest insert err ", err)
		return 0, err
	}
	return id, nil
}

func (p *Postgres) CreateRecommendation(ctx context.Context, recommendation models.Recommendation) (int64, error) {
	var id int64
	err := p.db.QueryRowContext(ctx, "INSERT INTO recommendations (consultationrequestid, text) VALUES ($1, $2) RETURNING ID", recommendation.ConsultationRequestID, recommendation.Text).Scan(&id)
	if err != nil {
		log.Println("CreateRecommendation insert err ", err)
		return 0, err
	}
	return id, nil
}

func (p *Postgres) GetRecommendationList(ctx context.Context, patientId int64) ([]models.PatientRecommendation, error) {
	var patientRecommendations []models.PatientRecommendation
	rows, err := p.db.QueryContext(ctx, "SELECT p.firstname, p.middlename, p.lastname, c.text as request_text, STRING_AGG (r.text, '; ') as recomendations FROM patients p JOIN consultationrequest c ON c.patientid=p.id JOIN recommendations r ON r.consultationrequestid = c.id WHERE p.id = $1 GROUP BY p.id, c.id", patientId)
	if err != nil {
		return []models.PatientRecommendation{}, err
	}
	for rows.Next() {
		var recomdendation models.PatientRecommendation
		err := rows.Scan(&recomdendation.FirstName, &recomdendation.MiddleName, &recomdendation.LastName, &recomdendation.RequestText, &recomdendation.Recommendations)
		if err != nil {
			log.Println("GetRecommendationList scan err ", err)
			continue
		}
		patientRecommendations = append(patientRecommendations, recomdendation)
	}
	return patientRecommendations, nil
}
