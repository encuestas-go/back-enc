package domain

type SocioeconomicStatus struct {
	IDUserType          int    `json:"id_user_type,omitempty"`
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