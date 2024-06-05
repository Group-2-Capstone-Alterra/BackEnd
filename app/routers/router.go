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

	_order_ProductData "PetPalApp/features/order_product/data"
	_order_ProductHandler "PetPalApp/features/order_product/handler"
	_order_ProductService "PetPalApp/features/order_product/service"

	_transactionData "PetPalApp/features/transaction/data"
	_transactionHandler "PetPalApp/features/transaction/handler"
	_transactionService "PetPalApp/features/transaction/service"

	_paymentData "PetPalApp/features/payment/data"
	_paymentHandler "PetPalApp/features/payment/handler"
	_paymentService "PetPalApp/features/payment/service"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB, s3 *s3.S3, s3Bucket string) {
	hashService := encrypts.NewHashService()
	helperUserService := helperuser.NewHelperService()
	helperService := helper.NewHelperService(s3, s3Bucket, _adminData.New(db), _userData.New(db, helperUserService))

	userData := _userData.New(db, helperUserService)
	userService := _userService.New(userData, hashService, helperService)
	userHandlerAPI := _userHandler.New(userService, hashService)

	productData := _productData.New(db, helperService)
	productService := _productService.New(productData, helperService)
	productHandlerAPI := _productHandler.New(productService, helperService)

	doctorData := _doctorData.New(db)
	doctorService := _doctorService.New(doctorData, helperService)
	doctorHandlerAPI := _doctorHandler.New(doctorService)

	adminData := _adminData.New(db)
	adminService := _adminService.New(adminData, hashService, doctorData, helperService)
	adminHandlerAPI := _adminHandler.New(adminService)

	orderData := _orderData.New(db)
	orderService := _orderService.New(orderData)
	orderHandlerAPI := _orderHandler.New(orderService)

	consultationData := _consultationData.New(db)
	consultationService := _consultationService.New(consultationData, doctorData, adminData)
	consultationHandlerAPI := _consultationHandler.New(consultationService, userData, doctorData)

	chatData := _chatData.New(db)
	chatService := _chatService.New(chatData, consultationData, doctorData, userData, adminData)
	chatHandlerAPI := _chatHandler.New(chatService, consultationData, userData, doctorData)

	order_ProductData := _order_ProductData.New(db)
	order_ProductService := _order_ProductService.New(order_ProductData)
	order_productHandlerAPI := _order_ProductHandler.New(order_ProductService)

	transactionData := _transactionData.New(db)
	transactionService := _transactionService.New(transactionData)
	transactionHandlerAPI := _transactionHandler.New(transactionService)

	paymentData := _paymentData.New(db)
	paymentService := _paymentService.New(paymentData)
	paymentHandlerAPI := _paymentHandler.New(paymentService)


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

	//chats
	e.POST("/chats/:id", chatHandlerAPI.CreateChat, middlewares.JWTMiddleware())
	e.GET("/chats/:id", chatHandlerAPI.GetChats, middlewares.JWTMiddleware())
	e.DELETE("/chats/:id", chatHandlerAPI.Delete, middlewares.JWTMiddleware())

	//orders
	e.POST("/orders", orderHandlerAPI.CreateOrder, middlewares.JWTMiddleware())
	e.GET("/orders", orderHandlerAPI.GetOrdersByUserID, middlewares.JWTMiddleware())

	//consultation
	e.POST("/consultations", consultationHandlerAPI.CreateConsultation, middlewares.JWTMiddleware())
	e.GET("/consultations", consultationHandlerAPI.GetConsultations, middlewares.JWTMiddleware())
	e.GET("/consultations/user", consultationHandlerAPI.GetConsultationsByUserID, middlewares.JWTMiddleware())
	e.GET("/consultations/doctor/:doctor_id", consultationHandlerAPI.GetConsultationsByDoctorID, middlewares.JWTMiddleware())
	e.PUT("/consultations/:consultation_id", consultationHandlerAPI.UpdateConsultationResponse, middlewares.JWTMiddleware())

	//order_products
	e.POST("/order_products", order_productHandlerAPI.CreateOrderProduct, middlewares.JWTMiddleware())
	e.GET("/order-products/:order_id", order_productHandlerAPI.GetOrderProductsByOrderID, middlewares.JWTMiddleware())

	//transactions
	e.POST("/transactions", transactionHandlerAPI.CreateTransaction, middlewares.JWTMiddleware())
	e.GET("/transactions/:user_id", transactionHandlerAPI.GetTransactionsByUserID, middlewares.JWTMiddleware())

	//payments
	e.POST("/payments", paymentHandlerAPI.CreatePayment, middlewares.JWTMiddleware())
}
