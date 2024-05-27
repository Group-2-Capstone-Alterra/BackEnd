package configs

import (
	"os"
	"strconv"
)

var (
	JWT_SECRET string
	//ctx        = context.Background()
)

type AppConfig struct {

	// err := godotenv.Load() // Menambahkan pembacaan ENV file local
	// if err != nil {

	// 	panic("Error loading .env file")
	// }
	// app.VALIDATLOCALORSERVER = os.Getenv("VALIDATLOCALORSERVER")

	DB_USERNAME string
	DB_PASSWORD string
	DB_HOSTNAME string
	DB_PORT     int
	DB_NAME     string

	RD_ADDRESS  string
	RD_PASSWORD string
	RD_DATABASE int

	S3_BUCKET         string
	S3_REGION         string
	S3_ACCESKEY       string
	S3_SECRETACCESKEY string

	VALIDATLOCALORSERVER string
}

func ReadEnv() *AppConfig {
	var app = AppConfig{}

	app.VALIDATLOCALORSERVER = os.Getenv("VALIDATLOCALORSERVER")

	app.DB_USERNAME = os.Getenv("DBUSER")
	app.DB_PASSWORD = os.Getenv("DBPASS")
	app.DB_HOSTNAME = os.Getenv("DBHOST")
	portConv, errConv := strconv.Atoi(os.Getenv("DBPORT"))
	if errConv != nil {
		panic("error convert dbport")
	}
	app.DB_PORT = portConv
	app.DB_NAME = os.Getenv("DBNAME")
	JWT_SECRET = os.Getenv("JWTSECRET")

	//redis
	// app.RD_ADDRESS = os.Getenv("RDADDR")
	// app.RD_PASSWORD = os.Getenv("RDPASS")
	// dbConv, errConv := strconv.Atoi(os.Getenv("RDDB"))
	// if errConv != nil {
	// 	panic("error convert redis database")
	// }
	// app.RD_DATABASE = dbConv

	//s3
	app.S3_BUCKET = os.Getenv("S3BUCKETNAME")
	app.S3_REGION = os.Getenv("S3REGION")
	app.S3_ACCESKEY = os.Getenv("S3ACCESKEY")
	app.S3_SECRETACCESKEY = os.Getenv("S3SECRETACCESKEY")

	return &app
}

func InitConfig() *AppConfig {
	return ReadEnv()
}
