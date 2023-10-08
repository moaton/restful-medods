package models

type Recommendation struct {
	ID                    int64  `json:"id"`
	ConsultationRequestID int64  `json:"consultation_request_id"`
	Text                  string `json:"text"`
}
