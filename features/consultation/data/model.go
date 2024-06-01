package data

import (
	"PetPalApp/features/consultation"

	"gorm.io/gorm"
)

type Consultation struct {
    gorm.Model
    UserID        uint   
    DoctorID      uint   
    Consultation  string 
    Response      string 
    Status        string 
}

func (c *Consultation) ToCore() consultation.ConsultationCore {
    return consultation.ConsultationCore{
        ID:           c.ID,
        UserID:       c.UserID,
        DoctorID:     c.DoctorID,
        Consultation: c.Consultation,
        Response:     c.Response,
        Status:       c.Status,
        CreatedAt:    c.CreatedAt,
    }
}