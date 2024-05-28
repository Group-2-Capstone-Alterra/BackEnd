package databases

import (
	"PetPalApp/app/configs"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func InitS3(cfg *configs.AppConfig) (*s3.S3, string) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(cfg.S3_REGION),
		Credentials: credentials.NewStaticCredentials(cfg.S3_ACCESKEY, cfg.S3_SECRETACCESKEY, ""),
	})
	if err != nil {
		log.Fatalf("Failed to create AWS session: %v", err)
	}
	s3Client := s3.New(sess)
	return s3Client, cfg.S3_BUCKET

}
