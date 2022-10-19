package dto

type UserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Birthdate string `json:"birthdate"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Document  string `json:"document"`
	City      string `json:"city"`
	Zipcode   string `json:"zipcode"`
	Phone     string `json:"phone"`
	Status    string `json:"status"`
}
