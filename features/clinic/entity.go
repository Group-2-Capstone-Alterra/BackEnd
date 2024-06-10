package clinic

import (
	"PetPalApp/features/availdaydoctor"
	"PetPalApp/features/servicedoctor"
)

type Core struct {
	ID                uint                `json:"admin_id,omitempty"`
	ClinicName        string              `json:"clinic_name,omitempty"`
	ClinicPicture     string              `json:"clinic_picture,omitempty"`
	Open              availdaydoctor.Core `json:"open,omitempty"`
	Service           servicedoctor.Core  `json:"service,omitempty"`
	Veterinary        string              `json:"veterinary,omitempty"`
	VeterinaryPicture string              `json:"veterinary_picture,omitempty"`
	About             string              `json:"about,omitempty"`
	Price             float32             `json:"price,omitempty"`
	Location          string              `json:"location,omitempty"`
	Coordinate        string              `json:"coordinate,omitempty"`
	Distance          float64             `json:"distance,omitempty"`
}
