package handler

//request
type AdminRequest struct {
	FullName           string `json:"full_name" form:"full_name"`
	Email              string `json:"email" form:"email"`
	NumberPhone        string `json:"number_phone" form:"number_phone"`
	Address            string `json:"address" form:"address"`
	Password           string `json:"password" form:"password"`
	KetikUlangPassword string `json:"ketik_ulang_password" form:"ketik_ulang_password"`
	ProfilePicture     string `json:"profile_picture" form:"profile_picture"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
