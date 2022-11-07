package dto

type UserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Birthdate string `json:"birthdate"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Document  string `json:"document"`
	Phone     string `json:"phone"`
	Status    string `json:"status"`
	IsAdmin   bool   `json:"is_admin"`
}

type UserRequest struct {
	Name      string `json:"name"`
	Birthdate string `json:"birthdate"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Document  string `json:"document"`
	Phone     string `json:"phone"`
	IsAdmin   bool   `json:"is_admin,omitempty"`
}
