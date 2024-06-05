package doctor

import (
	_avail "PetPalApp/features/availdaydoctor"
	_service "PetPalApp/features/servicedoctor"
	"io"
)

type Core struct {
	ID             uint
	AdminID        uint
	FullName       string
	About          string
	Price          float32
	ProfilePicture string
	AvailableDay   _avail.Core
	ServiceDoctor  _service.Core
}

type DoctorModel interface {
	AddDoctor(Core) error
	SelectByAdminId(id uint) (*Core, error)
	SelectDoctorById(id uint) (*Core, error)
	SelectAvailDayById(id uint) (*_avail.Core, error)
	SelectAllDoctor() ([]Core, error)
	PutByIdAdmin(AdminID uint, input Core) error
}

type DoctorService interface {
	AddDoctor(core Core, file io.Reader, handlerFilename string) (string, error)
	GetDoctorByIdAdmin(adminID uint) (*Core, error)
	GetAvailDoctorByIdDoctor(doctorID uint) (*_avail.Core, error)
	UpdateByIdAdmin(AdminId uint, input Core, file io.Reader, handlerFilename string) (string, error)
}
