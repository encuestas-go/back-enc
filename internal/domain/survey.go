package domain

type SocioeconomicStatus struct {
	IDUser              int    `json:"id_user,omitempty"`
	FullName            string `json:"full_name,omitempty"`
	BirthDate           string `json:"birth_date,omitempty"`
	Nationality         string `json:"nationality,omitempty"`
	Gender              string `json:"gender,omitempty"`
	Age                 int    `json:"age,omitempty"`
	MaritalStatus       string `json:"marital_status,omitempty"`
	ResidenceAddress    string `json:"residence_address,omitempty"`
	ResidenceCity       string `json:"residence_city,omitempty"`
	PostalCode          int    `json:"postal_code,omitempty"`
	State               string `json:"state,omitempty"`
	SocioeconomicStatus string `json:"socioeconomic_status,omitempty"`
	Language            string `json:"language,omitempty"`
	DegreeAspired       string `json:"degree_aspired,omitempty"`
	LastDegreeFather    string `json:"last_degree_father,omitempty"`
	LastDegreeMother    string `json:"last_degree_mother,omitempty"`
}

type EconomicStatus struct {
	IDUser                int     `json:"id_user,omitempty"`
	CurrentStatus         string  `json:"current_status,omitempty"`
	JobTitle              string  `json:"job_title,omitempty"`
	EmployerEstablishment string  `json:"employer_establishment,omitempty"`
	EmploymentType        string  `json:"employment_type,omitempty"`
	Salary                float64 `json:"salary,omitempty"`
	AmountType            string  `json:"amount_type,omitempty"`
	WorkBenefitsType      string  `json:"work_benefits_type,omitempty"`
}

type TransportManagement struct {
	UserID              int    `json:"user_id,omitempty"`
	PrimaryTransport    string `json:"primary_transport,omitempty"`
	SecondTransport     string `json:"second_transport,omitempty"`
	UsageFrequency      string `json:"usage_frequency,omitempty"`
	AccesiblePoints     bool   `json:"accesible_points,omitempty"`
	FrequentDestination string `json:"frequent_destination,omitempty"`
	TravelTime          string `json:"travel_time,omitempty"`
}
