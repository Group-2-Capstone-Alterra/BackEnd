package handler

import "PetPalApp/features/admin"

type AdminResponse struct {
	ID             uint   `json:"id"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	NumberPhone    string `json:"number_phone"`
	Role           string `json:"role"`
	Address        string `json:"address"`
	ProfilePicture string `json:"profile_picture"`
	Coordinate     string `json:"coordinate"`
}

type LoginResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func ResponseProfile(adminResponse admin.Core) AdminResponse {
	result := AdminResponse{
		ID:             adminResponse.ID,
		FullName:       adminResponse.FullName,
		Email:          adminResponse.Email,
		NumberPhone:    adminResponse.NumberPhone,
		Role:           adminResponse.Role,
		Address:        adminResponse.Address,
		ProfilePicture: adminResponse.ProfilePicture,
		Coordinate:     adminResponse.Coordinate,
	}
	return result
}

func ResponseLogin(adminResponse *admin.Core) LoginResponse {
	result := LoginResponse{
		ID:       adminResponse.ID,
		FullName: adminResponse.FullName,
		Email:    adminResponse.Email,
		Token:    adminResponse.Token,
	}
	return result
}
