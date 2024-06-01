package helper

import (
	"PetPalApp/features/admin"
	"PetPalApp/features/product"
	"PetPalApp/features/user"
	"bytes"
	"fmt"
	"io"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type HelperInterface interface {
	ConvertToNullableString(value string) *string
	UploadProfilePicture(file io.Reader, fileName string) (string, error)
	UploadProductPicture(file io.Reader, fileName string) (string, error)
	DereferenceString(s *string) string
	SortProductsByDistance(iduser uint, products []product.Core) []product.Core
}

type helper struct {
	s3       *s3.S3
	s3Bucket string
	admin    admin.AdminModel
	user     user.DataInterface
}

func NewHelperService(s3 *s3.S3, s3Bucket string, admin admin.AdminModel, user user.DataInterface) HelperInterface {
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

	_, err := u.s3.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(u.s3Bucket),
		Key:    aws.String("profilepicture/" + fileName),
		Body:   bytes.NewReader(buf.Bytes()),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/profilepicture/%s", u.s3Bucket, aws.StringValue(u.s3.Config.Region), fileName), err
}

func (u *helper) UploadProductPicture(file io.Reader, fileName string) (string, error) {
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		return "", err
	}

	_, err := u.s3.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(u.s3Bucket),
		Key:    aws.String("productpicture/" + fileName),
		Body:   bytes.NewReader(buf.Bytes()),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/productpicture/%s", u.s3Bucket, aws.StringValue(u.s3.Config.Region), fileName), err
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
