package handler

import (
	"PetPalApp/features/user"
)

type UserResponse struct {
	ID             uint   `json:"id,omitempty"`
	FullName       string `json:"full_name,omitempty"`
	Email          string `json:"email,omitempty"`
	NumberPhone    string `json:"number_phone,omitempty"`
	Role           string `json:"role,omitempty"`
	Address        string `json:"address,omitempty"`
	Password       string `json:"password,omitempty"`
	ProfilePicture string `json:"profile_picture,omitempty"`
	Coordinate     string `json:"coordinate,omitempty"`
	Token          string `json:"token,omitempty"`
}

type ConsulUserReponse struct {
	ID             uint   `json:"id,omitempty"`
	FullName       string `json:"full_name,omitempty"`
	ProfilePicture string `json:"profile_picture,omitempty"`
}

type OrderResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
}

func OrderCoreToResponse(core user.Core) OrderResponse {
	return OrderResponse{
		ID:       core.ID,
		FullName: core.FullName,
	}
}

func ConsulCoreToGorm(userGorm user.Core) ConsulUserReponse {
	userCore := ConsulUserReponse{
		ID:             userGorm.ID,
		FullName:       userGorm.FullName,
		ProfilePicture: userGorm.ProfilePicture,
	}
	return userCore
}

func ResponseProfile(userGorm user.Core) UserResponse {
	result := UserResponse{
		ID:             userGorm.ID,
		FullName:       userGorm.FullName,
		Email:          userGorm.Email,
		NumberPhone:    userGorm.NumberPhone,
		Role:           userGorm.Role,
		Address:        userGorm.Address,
		ProfilePicture: userGorm.ProfilePicture,
		Coordinate:     userGorm.Coordinate,
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
