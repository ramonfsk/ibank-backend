package dto

type AccountResponse struct {
	ID          string  `json:"id,omitempty"`
	UserID      string  `json:"user_id"`
	OpeningDate string  `json:"opening_date"`
	Agency      string  `json:"agency"`
	Number      string  `json:"number"`
	CheckDigit  string  `json:"check_digit"`
	PIN         string  `json:"pin"`
	Balance     float64 `json:"balance"`
	Status      string  `json:"status"`
}

type NewAccountResponse struct {
	ID            string `json:"id_account,omitempty"`
	Agency        string `json:"agency"`
	NumberAccount string `json:"number_account"`
	CheckDigit    string `json:"check_digit"`
}
