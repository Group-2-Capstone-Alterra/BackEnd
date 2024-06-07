package handler

import (
	"PetPalApp/features/admin"
	"PetPalApp/features/availdaydoctor"
	"PetPalApp/features/doctor"
	"PetPalApp/features/servicedoctor"
)

type AllClinicResponse struct {
	ID         uint                `json:"admin_id,omitempty"`
	ClinicName string              `json:"clinic_name,omitempty"`
	Open       availdaydoctor.Core `json:"open,omitempty"`
	Service    servicedoctor.Core  `json:"service,omitempty"`
	Veterinary string              `json:"veterinary,omitempty"`
	Location   string              `json:"location,omitempty"`
}

func ResponseAllClinic(admin admin.Core, doctor doctor.Core) AllClinicResponse {
	result := AllClinicResponse{
		ID:         admin.ID,
		ClinicName: admin.FullName,
		Open:       availdaydoctor.Core{},
		Service:    servicedoctor.Core{},
		Veterinary: doctor.FullName,
		Location:   admin.Address,
	}
	return result
}
