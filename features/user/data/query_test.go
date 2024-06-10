package data

import (
	"PetPalApp/features/user"
	"PetPalApp/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestUserQuery_Insert(t *testing.T) {
	dsn := "root:password@tcp(localhost:3306)/test_db"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.NoError(t, err)

	helperuserMock := &mocks.HelperuserInterface{}
	userQuery := New(db, helperuserMock)

	input := user.Core{
		FullName: "John Doe",
		Email:    "johndoe@example.com",
	}

	helperuserMock.On("ConvertToNullableString", input.FullName).Return(&input.FullName)
	helperuserMock.On("ConvertToNullableString", input.Email).Return(&input.Email)

	err = userQuery.Insert(input)
	assert.NoError(t, err)

	helperuserMock.AssertExpectations(t)
}

func TestUserQuerySelectByEmail(t *testing.T) {
	dsn := "root:password@tcp(localhost:3306)/test_db"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.NoError(t, err)

	helperuserMock := &mocks.HelperuserInterface{}
	userQuery := New(db, helperuserMock)

	email := "johndoe@example.com"
	userData := User{
		Email: &email,
	}

	db.Create(&userData)

	helperuserMock.On("DereferenceString", userData.Email).Return(*userData.Email)

	result, err := userQuery.SelectByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, email, result.Email)

	helperuserMock.AssertExpectations(t)
}

func TestUserQuerySelectById(t *testing.T) {
	dsn := "root:password@tcp(localhost:3306)/test_db"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.NoError(t, err)

	helperuserMock := &mocks.HelperuserInterface{}
	userQuery := New(db, helperuserMock)

	id := uint(1)
	email := "johndoe@example.com"
	userData := User{
		Model: gorm.Model{ID: id},
		Email: &email,
	}

	db.Create(&userData)

	helperuserMock.On("DereferenceString", userData.Email).Return(*userData.Email)

	result, err := userQuery.SelectById(id)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, id, result.ID)

	helperuserMock.AssertExpectations(t)
}

func TestUserQueryPutById(t *testing.T) {
	dsn := "root:password@tcp(localhost:3306)/test_db"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.NoError(t, err)

	helperuserMock := &mocks.HelperuserInterface{}
	userQuery := New(db, helperuserMock)

	id := uint(1)
	email := "johndoe@example.com"
	userData := User{
		Model: gorm.Model{ID: id},
		Email: &email,
	}

	db.Create(&userData)

	input := user.Core{
		FullName: "Jane Doe",
		Email:    "janedoe@example.com",
	}

	helperuserMock.On("ConvertToNullableString", input.FullName).Return(&input.FullName)
	helperuserMock.On("ConvertToNullableString", input.Email).Return(&input.Email)

	err = userQuery.PutById(id, input)
	assert.NoError(t, err)

	var updatedUserData User
	db.First(&updatedUserData, id)
	assert.Equal(t, input.FullName, *updatedUserData.FullName)
	assert.Equal(t, input.Email, *updatedUserData.Email)

	helperuserMock.AssertExpectations(t)
}

func TestUserQueryDelete(t *testing.T) {
	dsn := "root:password@tcp(localhost:3306)/test_db"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	assert.NoError(t, err)

	helperuserMock := &mocks.HelperuserInterface{}
	userQuery := New(db, helperuserMock)

	id := uint(1)
	email := "johndoe@example.com"
	userData := User{
		Model: gorm.Model{ID: id},
		Email: &email,
	}

	db.Create(&userData)

	err = userQuery.Delete(id)
	assert.NoError(t, err)

	var deletedUserData User
	db.First(&deletedUserData, id)
	assert.Nil(t, deletedUserData)

	helperuserMock.AssertExpectations(t)
}
