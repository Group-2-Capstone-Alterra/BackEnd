package doctor

type Core struct {
	FullName       string
	Email          string
	Specialization string
}

type DoctorModel interface {
	AddDoctor(Core) error
}

type DoctorService interface {
	AddDoctor(Core) error
}
