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

func ResponseProfile(userGorm user.Core) UserResponse {
	result := UserResponse{
		ID:             userGorm.ID,
		FullName:       userGorm.FullName,
		Email:          userGorm.Email,
		NumberPhone:    userGorm.NumberPhone,
		Address:        userGorm.Address,
		ProfilePicture: userGorm.ProfilePicture,
	}
	return result
}

func CoreToGorm(userGorm user.Core) UserResponse {
	userCore := UserResponse{
		ID:             userGorm.ID,
		FullName:       userGorm.FullName,
		Email:          userGorm.Email,
		NumberPhone:    userGorm.NumberPhone,
		Address:        userGorm.Address,
		Password:       userGorm.Password,
		ProfilePicture: userGorm.ProfilePicture,
	}
	return userCore
}

func ResponseLogin(userResponse *user.Core) UserResponse {
	result := UserResponse{
		ID:       userResponse.ID,
		FullName: userResponse.FullName,
		Token:    userResponse.Token,
	}
	return result
}
