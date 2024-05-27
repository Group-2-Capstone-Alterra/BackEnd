package main

import (
	"PetPalApp/app/configs"
	"PetPalApp/app/databases"
	"PetPalApp/app/migrations"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := configs.InitConfig()
	dbMysql := databases.InitDBMysql(cfg)
	migrations.InitMigrations(dbMysql)
	// s3Client, s3Bucket := databases.InitS3(cfg)

	e := echo.New()

	// e.Use(middlewares.RemoveTrailingSlash)
	// e.Use(middleware.CORSWithConfig(middlewares.CORSConfig()))

	// routers.InitRouter(e, dbMysql, s3Client, cfg, s3Bucket)
	e.Logger.Fatal(e.Start(":8080"))
}
