package handler

type AdminResponse struct {
	ID             uint   `json:"id"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	NumberPhone    string `json:"number_phone"`
	Role           string `json:"role"`
	Address        string `json:"address"`
	ProfilePicture string `json:"profile_picture"`
	Coordinate     string `json:"coordinate"`
}

type LoginResponse struct {
	ID       uint   `json:"id"`
	Role     string `json:"role"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}


