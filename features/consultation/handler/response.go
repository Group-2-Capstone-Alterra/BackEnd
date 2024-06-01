package handler

import "time"

type ConsultationResponse struct {
    ID           uint      `json:"id"`
    UserID       uint      `json:"user_id"`
    DoctorID     uint      `json:"doctor_id"`
    Consultation string    `json:"consultation"`
    Response     string    `json:"response"`
    Status       string    `json:"status"`
    CreatedAt    time.Time `json:"created_at"`
}
