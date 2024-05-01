package domain

type SocioeconomicStatus struct {
	ID                  int    `json:"id,omitempty"`
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
	ID                    int     `json:"id,omitempty"`
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
	ID                  int    `json:"id,omitempty"`
	UserID              int    `json:"user_id,omitempty"`
	PrimaryTransport    string `json:"primary_transport,omitempty"`
	SecondTransport     string `json:"second_transport,omitempty"`
	UsageFrequency      string `json:"usage_frequency,omitempty"`
	AccesiblePoints     bool   `json:"accesible_points,omitempty"`
	FrequentDestination string `json:"frequent_destination,omitempty"`
	TravelTime          string `json:"travel_time,omitempty"`
}

type HouseholdInfrastructure struct {
	ID                    int    `json:"id,omitempty"`
	UserID                int    `json:"user_id,omitempty"`
	Zone                  string `json:"zone,omitempty"`
	Permanence            string `json:"permanence,omitempty"`
	InfraestructureStatus string `json:"infraestructure_status,omitempty"`
	FloorType             string `json:"floor_type,omitempty"`
	RoofType              string `json:"roof_type,omitempty"`
	WallType              string `json:"wall_type,omitempty"`
	TotalMembers          int    `json:"total_members,omitempty"`
	TotalRooms            int    `json:"total_rooms,omitempty"`
	HouseholdEquipment    string `json:"household_equipment,omitempty"`
	BasicServices         string `json:"basic_services,omitempty"`
	OtherProperties       bool   `json:"other_properties,omitempty"`
}

type DemographicStatus struct {
	ID               int     `json:"id,omitempty"`
	UserID           int     `json:"user_id,omitempty"`
	HousingType      string  `json:"housing_type,omitempty"`
	HouseCondition   string  `json:"house_condition,omitempty"`
	OwnTransport     bool    `json:"own_transport,omitempty"`
	IncomeAmount     float64 `json:"income_amount,omitempty"`
	WorkingMembers   int     `json:"working_members,omitempty"`
	MembersUnderage  int     `json:"members_underage,omitempty"`
	MonthlyExpenses  float64 `json:"monthly_expenses,omitempty"`
	GovermentSupport bool    `json:"goverment_support,omitempty"`
}

type CulturalActivity struct {
	ID                   int    `json:"id,omitempty"`
	UserID               int    `json:"user_id,omitempty"`
	PreferredGame        string `json:"preferred_game,omitempty"`
	Hobby                string `json:"hobby,omitempty"`
	PreferredSport       string `json:"preferred_sport,omitempty"`
	ExerciseFrequency    string `json:"exercise_frequency,omitempty"`
	WorkshopType         string `json:"workshop_type,omitempty"`
	PreferredSocialEvent string `json:"preferred_social_event,omitempty"`
}

type Services struct {
	ID                 int    `json:"id,omitempty"`
	UserID             int    `json:"user_id,omitempty"`
	EnergyProvider     bool   `json:"energy_provider,omitempty"`
	WaterProvider      bool   `json:"water_provider,omitempty"`
	InternetProvider   string `json:"internet_provider,omitempty"`
	PhoneProvider      bool   `json:"phone_provider,omitempty"`
	TvProvider         bool   `json:"tv_provider,omitempty"`
	PaymentDueDate     string `json:"payment_due_date,omitempty"`
	AdditionalPayments string `json:"additional_payments,omitempty"`
	ServicesBill       int    `json:"services_bill,omitempty"`
}

type Event struct {
	ID          int    `json:"id,omitempty"`
	EventName   string `json:"event_name,omitempty"`
	Place       string `json:"place,omitempty"`
	Date        string `json:"date,omitempty"`
	Hour        string `json:"hour,omitempty"`
	Location    string `json:"location,omitempty"`
	Description string `json:"description,omitempty"`
	Category    string `json:"category,omitempty"`
	IDUser      int    `json:"id_user,omitempty"`
}
