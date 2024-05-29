package routers

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/utils/encrypts"
	"PetPalApp/utils/helper"

	_adminData "PetPalApp/features/admin/data"
	_adminHandler "PetPalApp/features/admin/handler"
	_adminService "PetPalApp/features/admin/service"

	_userData "PetPalApp/features/user/data"
	_userHandler "PetPalApp/features/user/handler"
	_userService "PetPalApp/features/user/service"

	_productData "PetPalApp/features/product/data"
	_productHandler "PetPalApp/features/product/handler"
	_productService "PetPalApp/features/product/service"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB, s3 *s3.S3, s3Bucket string) {
	hashService := encrypts.NewHashService()
	helperService := helper.NewHelperService(s3, s3Bucket)

	userData := _userData.New(db, helperService)
	userService := _userService.New(userData, hashService, helperService)
	userHandlerAPI := _userHandler.New(userService, hashService)

	productData := _productData.New(db)
	productService := _productService.New(productData, helperService)
	productHandlerAPI := _productHandler.New(productService)

	adminData := _adminData.New(db)
	adminService := _adminService.New(adminData, hashService)
	adminHandlerAPI := _adminHandler.New(adminService)

	//user
	e.POST("/users/create", userHandlerAPI.Register)
	e.POST("/users/login", userHandlerAPI.Login)
	e.GET("/users/profile", userHandlerAPI.Profile, middlewares.JWTMiddleware())
	e.PUT("/users/update", userHandlerAPI.UpdateUserById, middlewares.JWTMiddleware())

	//products
	e.POST("/products/insert", productHandlerAPI.AddProduct, middlewares.JWTMiddleware())

	//admins
	e.POST("/admin/register", adminHandlerAPI.Register)
	e.POST("/admin/login", adminHandlerAPI.Login)
}
