package dto

type AuthRequest struct {
	Token     string `json:"token"`
	UserID    string `json:"userID"`
	RouteName string `json:"routeName"`
}

type AuthResponse struct {
	IsAuthorized bool `json:"isAuthorized"`
}
