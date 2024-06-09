package main

import (
	"PetPalApp/app/configs"
	"PetPalApp/app/databases"
	"PetPalApp/app/midtrans"
	"PetPalApp/app/migrations"
	"PetPalApp/app/routers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := configs.InitConfig()
	dbMysql := databases.InitDBMysql(cfg)
	migrations.InitMigrations(dbMysql)
	s3Client, s3Bucket := databases.InitS3(cfg)
	MidtransClient := midtrans.GetMidtransClient(cfg)

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	routers.InitRouter(e, dbMysql, s3Client, s3Bucket, MidtransClient)
	e.Logger.Fatal(e.Start(":8080"))
}
