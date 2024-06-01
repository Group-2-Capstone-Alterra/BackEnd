package handler

type ConsultationRequest struct {
	DoctorID     uint   `json:"doctor_id" form:"doctor_id" validate:"required"`
	Consultation string `json:"consultation" form:"consultation" validate:"required"`
}
