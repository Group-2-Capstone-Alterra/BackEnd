package data_test

import (
	"PetPalApp/features/user"
	"PetPalApp/features/user/data"
	"PetPalApp/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupDatabase() (*gorm.DB, error) {
	dsn := "avnadmin:AVNS_PLhXEJqSXBwMp2iyh2y@tcp(bagasdb-bagas-76bb.d.aivencloud.com:13407)/defaultdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&data.User{})
	return db, nil
}

func TestInsert(t *testing.T) {
	db, err := setupDatabase()
	assert.NoError(t, err)

	mockHelper := new(mocks.HelperuserInterface)
	userModel := data.New(db, mockHelper)

	mockHelper.On("ConvertToNullableString", mock.Anything).Return(new(string))

	input := user.Core{
		FullName: "John Doe",
		Email:    "johndoe@example.com",
	}

	err = userModel.Insert(input)
	assert.NoError(t, err)

	var userData data.User
	err = db.First(&userData, "email = ?", input.Email).Error
	assert.NoError(t, err)
	assert.Equal(t, *userData.Email, input.Email)
}

func TestSelectByEmail(t *testing.T) {
	db, err := setupDatabase()
	assert.NoError(t, err)

	mockHelper := new(mocks.HelperuserInterface)
	userModel := data.New(db, mockHelper)

	mockHelper.On("DereferenceString", mock.Anything).Return("")

	userData := data.User{
		FullName: new(string),
		Email:    new(string),
	}
	*userData.Email = "johndoe@example.com"
	db.Create(&userData)

	coreUser, err := userModel.SelectByEmail("johndoe@example.com")
	assert.NoError(t, err)
	assert.NotNil(t, coreUser)
	assert.Equal(t, coreUser.Email, "johndoe@example.com")
}

func TestSelectById(t *testing.T) {
	db, err := setupDatabase()
	assert.NoError(t, err)

	mockHelper := new(mocks.HelperuserInterface)
	userModel := data.New(db, mockHelper)

	mockHelper.On("DereferenceString", mock.Anything).Return("")

	userData := data.User{
		FullName: new(string),
		Email:    new(string),
	}
	*userData.Email = "johndoe@example.com"
	db.Create(&userData)

	coreUser, err := userModel.SelectById(userData.ID)
	assert.NoError(t, err)
	assert.NotNil(t, coreUser)
	assert.Equal(t, coreUser.Email, "johndoe@example.com")
}

func TestPutById(t *testing.T) {
	db, err := setupDatabase()
	assert.NoError(t, err)

	mockHelper := new(mocks.HelperuserInterface)
	userModel := data.New(db, mockHelper)

	mockHelper.On("ConvertToNullableString", mock.Anything).Return(new(string))

	userData := data.User{
		FullName: new(string),
		Email:    new(string),
	}
	*userData.Email = "johndoe@example.com"
	db.Create(&userData)

	updateData := user.Core{
		FullName: "Jane Doe",
		Email:    "janedoe@example.com",
	}

	err = userModel.PutById(userData.ID, updateData)
	assert.NoError(t, err)

	var updatedUser data.User
	err = db.First(&updatedUser, userData.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, *updatedUser.Email, "janedoe@example.com")
}

func TestDelete(t *testing.T) {
	db, err := setupDatabase()
	assert.NoError(t, err)

	mockHelper := new(mocks.HelperuserInterface)
	userModel := data.New(db, mockHelper)

	userData := data.User{
		FullName: new(string),
		Email:    new(string),
	}
	*userData.Email = "johndoe@example.com"
	db.Create(&userData)

	err = userModel.Delete(userData.ID)
	assert.NoError(t, err)

	var deletedUser data.User
	err = db.First(&deletedUser, userData.ID).Error
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}
