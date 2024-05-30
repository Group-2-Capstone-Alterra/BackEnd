package service

import (
	"PetPalApp/features/doctor"
	"errors"
)

type DoctorService struct {
    DoctorModel doctor.DoctorModel
}

func New(dm doctor.DoctorModel) doctor.DoctorService {
    return &DoctorService{
        DoctorModel: dm,
    }
}

func (ds *DoctorService) AddDoctor(doctor doctor.Core) error {
    if doctor.FullName == "" || doctor.Email == "" || doctor.Specialization == "" {
        return errors.New("[validation] Fullname/email/specialization tidak boleh kosong")
    }
    return ds.DoctorModel.AddDoctor(doctor)
}
