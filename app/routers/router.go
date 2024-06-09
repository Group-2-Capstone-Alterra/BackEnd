package routers

import (
	"PetPalApp/app/middlewares"
	"PetPalApp/utils/encrypts"
	"PetPalApp/utils/helper"
	"PetPalApp/utils/helperuser"

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

	_consultationData "PetPalApp/features/consultation/data"
	_consultationHandler "PetPalApp/features/consultation/handler"
	_consultationService "PetPalApp/features/consultation/service"

	_paymentData "PetPalApp/features/payment/data"
	_paymentHandler "PetPalApp/features/payment/handler"
	_paymentService "PetPalApp/features/payment/service"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo/v4"
	"github.com/veritrans/go-midtrans"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB, s3 *s3.S3, s3Bucket string, midtrans midtrans.Client) {
	hashService := encrypts.NewHashService()
	helperUserService := helperuser.NewHelperService()
	helperService := helper.NewHelperService(s3, s3Bucket, _adminData.New(db), _userData.New(db, helperUserService))

	userData := _userData.New(db, helperUserService)
	userService := _userService.New(userData, hashService, helperService)
	userHandlerAPI := _userHandler.New(userService, hashService)

	productData := _productData.New(db, helperService)
	productService := _productService.New(productData, helperService, &_adminData.AdminModel{})
	productHandlerAPI := _productHandler.New(productService, helperService)

	doctorData := _doctorData.New(db)
	doctorService := _doctorService.New(doctorData, helperService)
	doctorHandlerAPI := _doctorHandler.New(doctorService, doctorData)

	adminData := _adminData.New(db)
	adminService := _adminService.New(adminData, hashService, doctorData, helperService)
	adminHandlerAPI := _adminHandler.New(adminService, helperService)

	orderData := _orderData.New(db)
	orderService := _orderService.New(orderData)
	orderHandlerAPI := _orderHandler.New(orderService)

	consultationData := _consultationData.New(db)
	consultationService := _consultationService.New(consultationData, doctorData, adminData)
	consultationHandlerAPI := _consultationHandler.New(consultationService, userData, doctorData)

	chatData := _chatData.New(db)
	chatService := _chatService.New(chatData, consultationData, doctorData, userData, adminData)
	chatHandlerAPI := _chatHandler.New(chatService, consultationData, userData, doctorData, adminData)

	paymentData := _paymentData.New(db)
	paymentService := _paymentService.New(paymentData)
	paymentHandlerAPI := _paymentHandler.New(paymentService, midtrans)

	//user
	e.POST("/users/register", userHandlerAPI.Register)
	e.POST("/users/login", userHandlerAPI.Login)
	e.GET("/users/profile", userHandlerAPI.Profile, middlewares.JWTMiddleware())
	e.PATCH("/users/profile", userHandlerAPI.UpdateUserById, middlewares.JWTMiddleware())
	e.DELETE("/users", userHandlerAPI.Delete, middlewares.JWTMiddleware())

	//products
	e.POST("/products", productHandlerAPI.AddProduct, middlewares.JWTMiddleware())
	e.GET("/products", productHandlerAPI.GetAllProduct)
	e.GET("/products/:id", productHandlerAPI.GetProductById)
	e.PATCH("/products/:id", productHandlerAPI.UpdateProductById, middlewares.JWTMiddleware())
	e.DELETE("/products/:id", productHandlerAPI.Delete, middlewares.JWTMiddleware())

	//admins
	e.POST("/admins/register", adminHandlerAPI.Register)
	e.POST("/admins/login", adminHandlerAPI.Login)
	e.GET("/admins", adminHandlerAPI.GetProfile, middlewares.JWTMiddleware())
	e.DELETE("/admins", adminHandlerAPI.Delete, middlewares.JWTMiddleware())
	e.PUT("/admins", adminHandlerAPI.Update, middlewares.JWTMiddleware())
	e.GET("/clinics", adminHandlerAPI.GetAllClinic)

	//doctors
	e.POST("/doctors", doctorHandlerAPI.AddDoctor, middlewares.JWTMiddleware())
	e.GET("/doctors", doctorHandlerAPI.ProfileDoctor, middlewares.JWTMiddleware())
	e.PATCH("/doctors", doctorHandlerAPI.UpdateProfile, middlewares.JWTMiddleware())
	e.PATCH("/doctors/uploadpicture", doctorHandlerAPI.UploadDoctorPicture, middlewares.JWTMiddleware())
	e.DELETE("/doctors", doctorHandlerAPI.Delete, middlewares.JWTMiddleware())

	//chats
	e.POST("/chats/:id", chatHandlerAPI.CreateChat, middlewares.JWTMiddleware())
	e.GET("/chats/:id", chatHandlerAPI.GetChats, middlewares.JWTMiddleware())
	e.DELETE("/chats/:id", chatHandlerAPI.Delete, middlewares.JWTMiddleware())

	//orders
	e.POST("/orders", orderHandlerAPI.CreateOrder, middlewares.JWTMiddleware())
	e.GET("/orders", orderHandlerAPI.GetOrdersByUserID, middlewares.JWTMiddleware())

	//consultation
	e.POST("/consultations/:id", consultationHandlerAPI.CreateConsultation, middlewares.JWTMiddleware())
	e.GET("/consultations", consultationHandlerAPI.GetConsultations, middlewares.JWTMiddleware())
	e.GET("/consultations/user", consultationHandlerAPI.GetConsultationsByUserID, middlewares.JWTMiddleware())
	e.GET("/consultations/doctor/:doctor_id", consultationHandlerAPI.GetConsultationsByDoctorID, middlewares.JWTMiddleware())
	e.PATCH("/consultations/:consultation_id", consultationHandlerAPI.UpdateConsultationResponse, middlewares.JWTMiddleware())

	//payments
	e.POST("/payments", paymentHandlerAPI.CreatePayment, middlewares.JWTMiddleware())
}
