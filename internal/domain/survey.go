package domain

type SocioeconomicStatus struct {
	ID                  int     `json:"id,omitempty"`
	IDUser              int     `json:"id_user,omitempty"`
	FullName            string  `json:"full_name,omitempty"`
	BirthDate           string  `json:"birth_date,omitempty"`
	Nationality         string  `json:"nationality,omitempty"`
	Gender              string  `json:"gender,omitempty"`
	Age                 int     `json:"age,omitempty"`
	MaritalStatus       string  `json:"marital_status,omitempty"`
	Longitude           float64 `json:"longitude,omitempty"`
	Latitude            float64 `json:"latitude,omitempty"`
	ResidenceAddress    string  `json:"residence_address,omitempty"`
	SocioeconomicStatus string  `json:"socioeconomic_status,omitempty"`
	Language            string  `json:"language,omitempty"`
	DegreeAspired       string  `json:"degree_aspired,omitempty"`
	LastDegreeFather    string  `json:"last_degree_father,omitempty"`
	LastDegreeMother    string  `json:"last_degree_mother,omitempty"`
}

type UserAddressMap struct {
	IDUser           int     `json:"id_user,omitempty"`
	CompleteAddress  string  `json:"complete_address,omitempty"`
	LongitudeAddress float64 `json:"longitude_address,omitempty"`
	LatitudeAddress  float64 `json:"latitude_address,omitempty"`
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
	Date                  string  `json:"date,omitempty"`
}

type TransportManagement struct {
	ID                  int    `json:"id,omitempty"`
	UserID              int    `json:"user_id,omitempty"`
	PrimaryTransport    string `json:"primary_transport,omitempty"`
	SecondTransport     string `json:"second_transport,omitempty"`
	UsageFrequency      string `json:"usage_frequency,omitempty"`
	AccesiblePoints     bool   `json:"accesible_points"`
	FrequentDestination string `json:"frequent_destination,omitempty"`
	TravelTime          string `json:"travel_time,omitempty"`
	Date                string `json:"date,omitempty"`
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
	OtherProperties       bool   `json:"other_properties"`
}

type DemographicStatus struct {
	ID               int     `json:"id,omitempty"`
	UserID           int     `json:"user_id,omitempty"`
	HousingType      string  `json:"housing_type,omitempty"`
	HouseCondition   string  `json:"house_condition,omitempty"`
	OwnTransport     bool    `json:"own_transport"`
	IncomeAmount     float64 `json:"income_amount,omitempty"`
	WorkingMembers   int     `json:"working_members,omitempty"`
	MembersUnderage  int     `json:"members_underage,omitempty"`
	MonthlyExpenses  float64 `json:"monthly_expenses,omitempty"`
	GovermentSupport bool    `json:"goverment_support"`
	Date             string  `json:"date,omitempty"`
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
	Date                 string `json:"date,omitempty"`
}

type Services struct {
	ID                 int     `json:"id,omitempty"`
	UserID             int     `json:"user_id,omitempty"`
	EnergyProvider     bool    `json:"energy_provider"`
	WaterProvider      bool    `json:"water_provider"`
	InternetProvider   string  `json:"internet_provider,omitempty"`
	PhoneProvider      bool    `json:"phone_provider"`
	TvProvider         bool    `json:"tv_provider"`
	PaymentDueDate     string  `json:"payment_due_date,omitempty"`
	AdditionalPayments string  `json:"additional_payments,omitempty"`
	ServicesBill       float64 `json:"services_bill,omitempty"`
	Date               string  `json:"date,omitempty"`
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

type SatisfactoryLikertSurvey struct {
	ID                  int `json:"id,omitempty"`
	FacilidadUso        int `json:"facilidad_uso,omitempty"`
	ClaridadInstruccion int `json:"claridad_instruccion,omitempty"`
	RelevanciaContenido int `json:"relevancia_contenido,omitempty"`
	RapidezCarga        int `json:"rapidez_carga,omitempty"`
	SatisfaccionGeneral int `json:"satisfaccion_general,omitempty"`
	IDUsuario           int `json:"id_user,omitempty"`
}

type SatisfactorySurvey struct {
	ID            int    `json:"id,omitempty"`
	IDUser        int    `json:"id_user,omitempty"`
	ScheduledDate string `json:"scheduled_date,omitempty"`
}

type AnswerResponseForum struct {
	Questions []Question `json:"questions,omitempty"`
}

type Question struct {
	ID           int      `json:"id,omitempty"`
	IDUser       int      `json:"id_user,omitempty"`
	Name         string   `json:"name,omitempty"`
	QuestionText string   `json:"question_text,omitempty"`
	Answers      []Answer `json:"answers,omitempty"`
}

type Answer struct {
	ID         int    `json:"id,omitempty"`
	IDUser     int    `json:"id_user,omitempty"`
	QuestionID int    `json:"question_id,omitempty"`
	Name       string `json:"name,omitempty"`
	AnswerText string `json:"answer_text,omitempty"`
}

// STRUCTS FOR CORRESPONDING REPORTS THAT ARE REQUIRED
type MostUsedTransportReport struct {
	PrimaryTransport string `json:"primary_transport,omitempty"`
	Quantity         int    `json:"quantity,omitempty"`
}

type IncomeAmountReport struct {
	IncomeAmount float64 `json:"income_amount,omitempty"`
	Quantity     int     `json:"quantity,omitempty"`
}

type HouseTypeAndConditionReport struct {
	HousingType    string `json:"housing_type,omitempty"`
	HouseCondition string `json:"house_condition,omitempty"`
	Quantity       int    `json:"quantity,omitempty"`
}

type InternetProviderReport struct {
	InternetProvider string `json:"internet_provider,omitempty"`
	Quantity         int    `json:"quantity,omitempty"`
}

type CulturalActivitiesReport struct {
	Dance             int `json:"dance,omitempty"`
	PlayInstrument    int `json:"play_instrument,omitempty"`
	Paint             int `json:"paint,omitempty"`
	Draw              int `json:"draw,omitempty"`
	DoExercise        int `json:"do_exercise,omitempty"`
	Read              int `json:"read,omitempty"`
	GoWalking         int `json:"go_walking,omitempty"`
	Movies            int `json:"movies,omitempty"`
	OtherActivities   int `json:"other_activities,omitempty"`
	Festivals         int `json:"festivals,omitempty"`
	Concerts          int `json:"concerts,omitempty"`
	ArtExposition     int `json:"art_exposition,omitempty"`
	LiteraturePoetry  int `json:"literature_poetry,omitempty"`
	Dances            int `json:"dances,omitempty"`
	Conferences       int `json:"conferences,omitempty"`
	RecreationalParks int `json:"recreational_parks,omitempty"`
	OtherEvents       int `json:"other_events,omitempty"`
}

type StudentSituationReport struct {
	CurrentSituation string `json:"current_situation,omitempty"`
	Quantity         int    `json:"quantity,omitempty"`
}
