package handler

type AdminResponse struct {
	FullName       string `json:"fullName"`
	Email          string `json:"email"`
	NumberPhone    string `json:"number_phone"`
	Address        string `json:"address"`
	ProfilePicture string `json:"profile_picture"`
}