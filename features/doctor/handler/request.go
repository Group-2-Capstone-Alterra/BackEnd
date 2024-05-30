package handler

type DoctorRequest struct {
	FullName       string `json:"fullname" form:"fullname"`
	Email          string `json:"email" form:"email"`
	Specialization string `json:"specialization" form:"specialization"`
}
