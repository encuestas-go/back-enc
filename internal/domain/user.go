package domain

// User domain contains all the informacion from a user
type User struct {
	Name        string `json:"name,omitempty"`
	MiddleName  string `json:"middle_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password"`
	IDUserType  int    `json:"id_user_type,omitempty"`
}

type UserLogin struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}
