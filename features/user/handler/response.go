package handler

import (
	"PetPalApp/features/user"
)

type UserResponse struct {
	ID             uint   `json:"id,omitempty"`
	FullName       string `json:"FullName,omitempty"`
	Email          string `json:"email,omitempty"`
	NumberPhone    string `json:"tanggal_lahir,omitempty"`
	Address        string `json:"foto,omitempty"`
	Password       string `json:"Password,omitempty"`
	ProfilePicture string `json:"ProfilePicture,omitempty"`
	Token          string `json:"token,omitempty"`
}

func ResponseProfile(userResponse *user.Core) UserResponse {
	result := UserResponse{
		ID:             userResponse.ID,
		FullName:       userResponse.FullName,
		Email:          userResponse.Email,
		NumberPhone:    userResponse.NumberPhone,
		Address:        userResponse.Address,
		Password:       userResponse.Password,
		ProfilePicture: userResponse.ProfilePicture,
	}
	return result
}

// func CoreToGorm(userGorm user.Core) UserResponse {
// 	userCore := UserResponse{
// 		ID:           userGorm.ID,
// 		Nama:         userGorm.Nama,
// 		Email:        userGorm.Email,
// 		TanggalLahir: userGorm.TanggalLahir,
// 		Foto:         userGorm.Foto,
// 	}

// 	return userCore
// }

func ResponseLogin(userResponse *user.Core) UserResponse {
	result := UserResponse{
		ID:       userResponse.ID,
		FullName: userResponse.FullName,
		Token:    userResponse.Token,
	}
	return result
}
