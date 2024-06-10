package data_test

import (
	"PetPalApp/features/user/data"
	"PetPalApp/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
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
