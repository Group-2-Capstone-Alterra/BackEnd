package data

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama         string `json:"nama" form:"nama"`
	Email        string `gorm:"unique" json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	TanggalLahir string `json:"tanggal_lahir" form:"tanggal_lahir"`
	Foto         string `json:"foto" form:"foto"`
}
