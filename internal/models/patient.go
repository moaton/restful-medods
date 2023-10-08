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

type ActiveIngredient struct {
	Name     string `json:"name"`
	Strength string `json:"strength"`
}

type Medicine struct {
	ProductNdc        string             `json:"product_ndc"`
	BrandName         string             `json:"brand_name"`
	ActiveIngredients []ActiveIngredient `json:"active_ingredients"`
	Route             []string           `json:"route"`
	LabelerName       string             `json:"labeler_name"`
}

type PatientRecommendation struct {
	FirstName       string     `json:"first_name"`
	MiddleName      string     `json:"middle_name"`
	LastName        string     `json:"last_name"`
	RequestText     string     `json:"request_text"`
	Recommendations string     `json:"recommendations"`
	Medicines       []Medicine `json:"medicine"`
}

type MedicineResponse struct {
	Results []Medicine `json:"results"`
}
