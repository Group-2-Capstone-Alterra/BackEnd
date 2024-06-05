package handler

import (
	"PetPalApp/features/user"
)

type UserResponse struct {
	ID             uint   `json:"id,omitempty"`
	FullName       string `json:"full_name,omitempty"`
	Email          string `json:"email,omitempty"`
	NumberPhone    string `json:"number_phone,omitempty"`
	Address        string `json:"address,omitempty"`
	Password       string `json:"password,omitempty"`
	ProfilePicture string `json:"profile_picture,omitempty"`
	Coordinate     string `json:"coordinate,omitempty"`
	Role           string `json:"role,omitempty"`
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
		Coordinate:     userGorm.Coordinate,
		Role:           userGorm.Role,
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
		Coordinate:     userGorm.Coordinate,
	}
	return userCore
}

func ResponseLogin(userResponse *user.Core) UserResponse {
	result := UserResponse{
		ID:       userResponse.ID,
		FullName: userResponse.FullName,
		Role:     userResponse.Role,
		Token:    userResponse.Token,
	}
	return result
}
