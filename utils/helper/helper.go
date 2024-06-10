package helper

import (
	"PetPalApp/features/admin"
	"PetPalApp/features/clinic"
	"PetPalApp/features/product"
	"PetPalApp/features/user"
	"bytes"
	"fmt"
	"io"
	"math"
	"mime"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type HelperInterface interface {
	ConvertToNullableString(value string) *string
	UploadProfilePicture(file io.Reader, fileName string) (string, error)
	UploadAdminPicture(file io.Reader, fileName string) (string, error)
	UploadProductPicture(file io.Reader, fileName string) (string, error)
	UploadDoctorPicture(file io.Reader, fileName string) (string, error)
	DereferenceString(s *string) string
	SortProductsByDistance(iduser uint, products []product.Core) []product.Core
	SortClinicsByDistance(userid uint, clnics []clinic.Core) []clinic.Core
}

type helper struct {
	s3       *s3.S3
	s3Bucket string
	admin    admin.AdminModel
	user     user.UserModel
}

func NewHelperService(s3 *s3.S3, s3Bucket string, admin admin.AdminModel, user user.UserModel) HelperInterface {
	return &helper{
		s3:       s3,
		s3Bucket: s3Bucket,
		admin:    admin,
		user:     user,
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

	// Determine the content type based on the file extension
	contentType := mime.TypeByExtension(filepath.Ext(fileName))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	_, err := u.s3.PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(u.s3Bucket),
		Key:                aws.String("profilepicture/" + fileName),
		Body:               bytes.NewReader(buf.Bytes()),
		ACL:                aws.String("public-read"),
		ContentType:        aws.String(contentType),
		ContentDisposition: aws.String("inline; filename=\"" + fileName + "\""),
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/profilepicture/%s", u.s3Bucket, aws.StringValue(u.s3.Config.Region), fileName), err
}

func (u *helper) UploadAdminPicture(file io.Reader, fileName string) (string, error) {
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		return "", err
	}

	// Determine the content type based on the file extension
	contentType := mime.TypeByExtension(filepath.Ext(fileName))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	_, err := u.s3.PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(u.s3Bucket),
		Key:                aws.String("adminpicture/" + fileName),
		Body:               bytes.NewReader(buf.Bytes()),
		ACL:                aws.String("public-read"),
		ContentType:        aws.String(contentType),
		ContentDisposition: aws.String("inline; filename=\"" + fileName + "\""),
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/adminpicture/%s", u.s3Bucket, aws.StringValue(u.s3.Config.Region), fileName), err
}

func (u *helper) UploadProductPicture(file io.Reader, fileName string) (string, error) {
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		return "", err
	}

	// Determine the content type based on the file extension
	contentType := mime.TypeByExtension(filepath.Ext(fileName))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	_, err := u.s3.PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(u.s3Bucket),
		Key:                aws.String("productpicture/" + fileName),
		Body:               bytes.NewReader(buf.Bytes()),
		ACL:                aws.String("public-read"),
		ContentType:        aws.String(contentType),
		ContentDisposition: aws.String("inline; filename=\"" + fileName + "\""),
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/productpicture/%s", u.s3Bucket, aws.StringValue(u.s3.Config.Region), fileName), err
}

func (u *helper) UploadDoctorPicture(file io.Reader, fileName string) (string, error) {
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		return "", err
	}

	// Determine the content type based on the file extension
	contentType := mime.TypeByExtension(filepath.Ext(fileName))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	_, err := u.s3.PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(u.s3Bucket),
		Key:                aws.String("doctorpicture/" + fileName),
		Body:               bytes.NewReader(buf.Bytes()),
		ACL:                aws.String("public-read"),
		ContentType:        aws.String(contentType),
		ContentDisposition: aws.String("inline; filename=\"" + fileName + "\""),
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/doctorpicture/%s", u.s3Bucket, aws.StringValue(u.s3.Config.Region), fileName), err
}

func (u *helper) SortProductsByDistance(userid uint, products []product.Core) []product.Core {

	user, _ := u.user.SelectById(userid)
	userCoorConv := strings.Split(user.Coordinate, ",")
	userLat, errUserLat := strconv.ParseFloat(userCoorConv[0], 64)
	if errUserLat != nil {
		return nil
	}
	userLong, errUserLong := strconv.ParseFloat(userCoorConv[1], 64)
	if errUserLong != nil {
		return nil
	}

	for i := range products {
		//get coordinate admin
		adminID := products[i].IdUser
		dataAdmin, _ := u.admin.AdminById(adminID)

		adminCoor := strings.Split(dataAdmin.Coordinate, ",")
		adminLat, errAdminLat := strconv.ParseFloat(adminCoor[0], 64)
		if errAdminLat != nil {
			return nil
		}
		adminLong, errAdminLong := strconv.ParseFloat(adminCoor[1], 64)
		if errAdminLong != nil {
			return nil
		}

		//get distance
		distance := Distance(userLat, userLong, adminLat, adminLong)
		products[i].Distance = distance
	}

	sort.Slice(products, func(i, j int) bool {
		return products[i].Distance < products[j].Distance
	})

	return products
}

func (u *helper) SortClinicsByDistance(userid uint, clinics []clinic.Core) []clinic.Core {
	user, _ := u.user.SelectById(userid)
	userCoorConv := strings.Split(user.Coordinate, ",")
	userLat, errUserLat := strconv.ParseFloat(userCoorConv[0], 64)
	if errUserLat != nil {
		return nil
	}
	userLong, errUserLong := strconv.ParseFloat(userCoorConv[1], 64)
	if errUserLong != nil {
		return nil
	}

	for i := range clinics {

		adminID := clinics[i].ID
		dataAdmin, _ := u.admin.AdminById(adminID)

		adminCoor := strings.Split(dataAdmin.Coordinate, ",")
		adminLat, errAdminLat := strconv.ParseFloat(adminCoor[0], 64)
		if errAdminLat != nil {
			return nil
		}

		adminLong, errAdminLong := strconv.ParseFloat(adminCoor[1], 64)
		if errAdminLong != nil {
			return nil
		}

		//get distance
		distance := Distance(userLat, userLong, adminLat, adminLong)
		clinics[i].Distance = distance
	}
	sort.Slice(clinics, func(i, j int) bool {
		return clinics[i].Distance < clinics[j].Distance
	})
	return clinics

}

func Distance(userLat, userLon, adminLat, adminLon float64) float64 {
	const R = 6371 // Radius of the Earth in km
	var dLat, dLon float64
	dLat = (adminLat - userLat) * math.Pi / 180
	dLon = (adminLon - userLon) * math.Pi / 180
	userLat = userLat * math.Pi / 180
	adminLat = adminLat * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(userLat)*math.Cos(adminLat)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}
