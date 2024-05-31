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

	_doctorData "PetPalApp/features/doctor/data"
	_doctorHandler "PetPalApp/features/doctor/handler"
	_doctorService "PetPalApp/features/doctor/service"

	_chatData "PetPalApp/features/chat/data"
	_chatHandler "PetPalApp/features/chat/handler"
	_chatService "PetPalApp/features/chat/service"

	_orderData "PetPalApp/features/order/data"
	_orderHandler "PetPalApp/features/order/handler"
	_orderService "PetPalApp/features/order/service"

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

	doctorData := _doctorData.New(db)
	doctorService := _doctorService.New(doctorData)
	doctorHandlerAPI := _doctorHandler.New(doctorService)

	chatData := _chatData.New(db)
	chatService := _chatService.New(chatData)
	chatHandlerAPI := _chatHandler.New(chatService)

	orderData := _orderData.New(db)
	orderService := _orderService.New(orderData)
	orderHandlerAPI := _orderHandler.New(orderService)

	//user
	e.POST("/users/register", userHandlerAPI.Register)
	e.POST("/users/login", userHandlerAPI.Login)
	e.GET("/users/profile", userHandlerAPI.Profile, middlewares.JWTMiddleware())
	e.PUT("/users/profile", userHandlerAPI.UpdateUserById, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandlerAPI.Delete, middlewares.JWTMiddleware())

	//products
	e.POST("/products/insert", productHandlerAPI.AddProduct, middlewares.JWTMiddleware())
	e.GET("/products", productHandlerAPI.GetAllProduct)
	e.GET("/products/:id", productHandlerAPI.GetProductById)
	e.PUT("/admins/products/:id", productHandlerAPI.UpdateProductById, middlewares.JWTMiddleware())

	//admins
	e.POST("/admins/register", adminHandlerAPI.Register)
	e.POST("/admins/login", adminHandlerAPI.Login)
	e.GET("/admins", adminHandlerAPI.GetProfile, middlewares.JWTMiddleware())
	e.DELETE("/admins", adminHandlerAPI.Delete, middlewares.JWTMiddleware())
	e.PUT("/admins", adminHandlerAPI.Update, middlewares.JWTMiddleware())

	//doctors
	e.POST("/doctors", doctorHandlerAPI.AddDoctor, middlewares.JWTMiddleware())

	//chats
	e.POST("/chats", chatHandlerAPI.CreateChat, middlewares.JWTMiddleware())
	e.GET("/chats", chatHandlerAPI.GetChats, middlewares.JWTMiddleware())

	//orders
	e.POST("/orders", orderHandlerAPI.CreateOrder, middlewares.JWTMiddleware())
	e.GET("/orders", orderHandlerAPI.GetOrdersByUserID, middlewares.JWTMiddleware())
}
