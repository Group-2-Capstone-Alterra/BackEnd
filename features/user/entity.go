package user

import (
	"time"
)

type Core struct {
	ID                 uint
	Nama               string
	Email              string
	Password           string
	KetikUlangPassword string
	TanggalLahir       string
	Foto               string
	Token              string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
