package models

type Patient struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	MiddleName  string `json:"middle_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	DateOfBirst string `json:"date_of_brith"`
}

type PatientRecommendation struct {
	FirstName       string `json:"first_name"`
	MiddleName      string `json:"middle_name"`
	LastName        string `json:"last_name"`
	RequestText     string `json:"request_text"`
	Recommendations string `json:"recommendations"`
}
