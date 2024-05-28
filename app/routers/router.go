package routers

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/utils/encrypts"
	"PetPalApp/utils/helper"

	_userData "PetPalApp/features/user/data"
	_userHandler "PetPalApp/features/user/handler"
	_userService "PetPalApp/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {
	hashService := encrypts.NewHashService()
	helperService := helper.NewHelperService()

	userData := _userData.New(db, helperService)
	userService := _userService.New(userData, hashService, helperService)
	userHandlerAPI := _userHandler.New(userService, hashService)

	e.POST("/user/register", userHandlerAPI.Register)
	e.POST("/user/login", userHandlerAPI.Login)
	e.GET("/user/profile", userHandlerAPI.Profile, middlewares.JWTMiddleware())
}
