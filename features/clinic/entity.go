package clinic

import (
	"PetPalApp/features/availdaydoctor"
	"PetPalApp/features/servicedoctor"
)

type Core struct {
	ID         uint                `json:"admin_id"`
	ClinicName string              `json:"clinic_name"`
	Open       availdaydoctor.Core `json:"open"`
	Service    servicedoctor.Core  `json:"service"`
	Veterinary string              `json:"veterinary"`
	Location   string              `json:"location"`
	Coordinate string              `json:"coordinate"`
	Distance   float64             `json:"distance"`
}
