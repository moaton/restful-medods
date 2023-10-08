package models

import "time"

type ConsultationRequest struct {
	ID        int64     `json:"id"`
	PatientID int64     `json:"patient_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
