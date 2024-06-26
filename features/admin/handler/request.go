package handler

type AdminRequest struct {
	FullName       string `json:"full_name" form:"full_name"`
	Email          string `json:"email" form:"email"`
	NumberPhone    string `json:"number_phone" form:"number_phone"`
	Address        string `json:"address" form:"address"`
	Password       string `json:"password" form:"password"`
	ProfilePicture string `json:"profile_picture" form:"profile_picture"`
	Coordinate     string `json:"coordinate" form:"coordinate"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
