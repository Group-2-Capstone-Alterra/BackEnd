package data

import (
	"PetPalApp/features/consultation"
	"time"

	"gorm.io/gorm"
)

type Consultation struct {
	gorm.Model
	UserID             uint
	DoctorID           uint
	Service            string
	ScheduledDate      string
	TransactionStatus  string `gorm:"default:'Pending'"`
	StatusConsultation string `gorm:"default:'New Consultation'"`
}

// Custom type for date-only format
type DateOnly struct {
	time.Time
}

// MarshalJSON to format the date when sending the response
func (d DateOnly) MarshalJSON() ([]byte, error) {
	return []byte(d.Time.Format(`"2006-01-02"`)), nil
}

// UnmarshalJSON to parse the date from the request
func (d *DateOnly) UnmarshalJSON(b []byte) error {
	var err error
	d.Time, err = time.Parse(`"2006-01-02"`, string(b))
	return err
}

func ToCore(c Consultation) consultation.ConsultationCore {
	return consultation.ConsultationCore{
		ID:                 c.ID,
		UserID:             c.UserID,
		DoctorID:           c.DoctorID,
		Service:            c.Service,
		TransactionStatus:  c.TransactionStatus,
		StatusConsultation: c.StatusConsultation,
		ScheduledDate:      c.ScheduledDate,
	}
}
