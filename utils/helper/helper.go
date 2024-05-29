package helper

import (
	"bytes"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type HelperInterface interface {
	ConvertToNullableString(value string) *string
	UploadProfilePicture(file io.Reader, fileName string) (string, error)
	DereferenceString(s *string) string
}

type helper struct {
	s3       *s3.S3
	s3Bucket string
}

func NewHelperService(s3 *s3.S3, s3Bucket string) HelperInterface {
	return &helper{
		s3:       s3,
		s3Bucket: s3Bucket,
	}
}

func (h *helper) ConvertToNullableString(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

// Fungsi helper untuk meng-dereference pointer string
func (h *helper) DereferenceString(s *string) string {
	if s == nil {
		return "" // atau nilai default lain yang sesuai
	}
	return *s
}

func (u *helper) UploadProfilePicture(file io.Reader, fileName string) (string, error) {
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		return "", err
	}

	_, err := u.s3.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(u.s3Bucket),
		Key:    aws.String("fotoprofile/" + fileName),
		Body:   bytes.NewReader(buf.Bytes()),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/profilepicture/%s", u.s3Bucket, aws.StringValue(u.s3.Config.Region), fileName), err
}
