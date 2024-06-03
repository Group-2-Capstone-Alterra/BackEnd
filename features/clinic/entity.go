package clinic

import (
	"PetPalApp/features/availdaydoctor"
)

type Core struct {
	ID         uint                `json:"admin_id"`
	ClinicName string              `json:"clinic_name"`
	Open       availdaydoctor.Core `json:"open"`
	Service    string              `json:"service"`
	Veterinary string              `json:"veterinary"`
	Location   string              `json:"location"`
	Coordinate string              `json:"coordinat"`
	Distance   float64             `json:"distance"`
}
