package handler

type DoctorResponse struct {
	FullName       string `json:"fullname"`
	Email          string `json:"email"`
	Specialization string `json:"specialization"`
}
