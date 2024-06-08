package service

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/features/user"
	"PetPalApp/utils/encrypts"
	"PetPalApp/utils/helper"
	"errors"
	"fmt"
	"io"
	"time"
)

type userService struct {
	userData      user.DataInterface
	hashService   encrypts.HashInterface
	helperService helper.HelperInterface
}

func New(ud user.DataInterface, hash encrypts.HashInterface, helper helper.HelperInterface) user.ServiceInterface {
	return &userService{
		userData:      ud,
		hashService:   hash,
		helperService: helper,
	}
}

const (
	errid = "ID must be a positive integer"
)

func (u *userService) Create(input user.Core) error {

	result, _ := u.hashService.HashPassword(input.Password)
	input.Password = result

	err := u.userData.Insert(input)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) Login(email string, password string) (data *user.Core, token string, err error) {
	data, err = u.userData.SelectByEmail(email)
	if err != nil {
		return nil, "", err
	}
	isLoginValid := u.hashService.CheckPasswordHash(data.Password, password)
	if !isLoginValid {
		return nil, "", errors.New("Invalid email or password. Please try again.")
	}
	token, errJWT := middlewares.CreateToken(int(data.ID), data.Role)
	if errJWT != nil {
		return nil, "", errJWT
	}
	return data, token, nil
}

func (u *userService) GetProfile(id uint) (data *user.Core, err error) {
	if id <= 0 {
		return nil, errors.New(errid)
	}
	return u.userData.SelectById(id)
}

func (u *userService) UpdateById(id uint, input user.Core, file io.Reader, handlerFilename string) (string, error) {
	if id <= 0 {
		return "", errors.New(errid)
	}

	if input.Password != "" {
		result, errHash := u.hashService.HashPassword(input.Password)
		if errHash != nil {
			return "", errHash
		}
		input.Password = result
	}
	if file != nil && handlerFilename != "" {
		timestamp := time.Now().Unix()
		fileName := fmt.Sprintf("%d_%s", timestamp, handlerFilename)
		photoFileName, errPhoto := u.helperService.UploadProfilePicture(file, fileName)
		if errPhoto != nil {
			return "", errPhoto
		}
		input.ProfilePicture = photoFileName
	}

	err := u.userData.PutById(id, input)
	if err != nil {
		return "", err
	}
	return input.ProfilePicture, nil
}

func (u *userService) Delete(id uint) error {
	if id <= 0 {
		return errors.New(errid)
	}
	return u.userData.Delete(id)
}
