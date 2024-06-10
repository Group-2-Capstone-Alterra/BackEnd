package service

import (
	"PetPalApp/features/admin"
	"PetPalApp/features/product"
	"PetPalApp/utils/helper"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type productService struct {
	productData product.ProductModel
	helper      helper.HelperInterface
	adminData   admin.AdminModel
}

func New(pd product.ProductModel, helper helper.HelperInterface, adminData admin.AdminModel) product.ProductService {
	return &productService{
		productData: pd,
		helper:      helper,
		adminData:   adminData,
	}
}

const (
	errid = "ID must be a positive integer"
)

func (p *productService) Create(id uint, input product.Core, file io.Reader, handlerFilename string) (string, error) {
	if id <= 0 {
		return "", echo.NewHTTPError(http.StatusBadRequest, errid)
	}
	if input.ProductName == "" || input.Price <= 0 || file == nil || input.Stock == 0 || input.Description == "" {
		return "", errors.New("Failed to add product. Please ensure all fields are filled in correctly.")
	}
	timestamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d_%s", timestamp, handlerFilename)
	photoFileName, errPhoto := p.helper.UploadProductPicture(file, fileName)
	if errPhoto != nil {
		return "", errPhoto
	}
	input.IdUser = id
	input.ProductPicture = photoFileName

	err := p.productData.Insert(input)
	if err != nil {
		return "", err
	}
	return input.ProductPicture, nil
}

func (p *productService) GetAll(userid uint, limit uint, role string, offset uint, sortStr string) ([]product.Core, error) {
	if role == "user" { // is user
		product, err := p.productData.SelectAll(limit, offset, sortStr)
		if err != nil {
			return nil, err
		}
		if sortStr == "lowest distance" || sortStr == "higest distance" {
			productSort := p.helper.SortProductsByDistance(userid, product)
			return productSort, nil
		} else {
			return product, nil
		}
	} else if role == "admin" { // admin
		product, err := p.productData.SelectAllAdmin(limit, userid, offset)
		if err != nil {
			return nil, err
		}
		return product, nil

	} else { //is guest
		product, err := p.productData.SelectAll(limit, offset, sortStr)
		if err != nil {
			return nil, err
		}
		return product, nil
	}
}

func (p *productService) GetProductById(id uint, userid uint) (data *product.Core, err error) {
	if id <= 0 {
		return nil, echo.NewHTTPError(http.StatusBadRequest, errid)
	}

	if userid != 0 {
		return p.productData.SelectByIdAdmin(id, userid)
	} else {
		return p.productData.SelectById(id)
	}
}

func (p *productService) GetProductByName(userid uint, limit uint, role string, offset uint, sortStr, name string) ([]product.Core, error) {
	if role == "user" { // is user
		product, err := p.productData.SelectByName(limit, offset, sortStr, name)
		if err != nil {
			return nil, err
		}
		if sortStr == "lowest distance" || sortStr == "higest distance" {
			productSort := p.helper.SortProductsByDistance(userid, product)
			return productSort, nil
		} else {
			return product, nil
		}
	} else if role == "admin" { // admin
		product, err := p.productData.SelectAllAdminByName(limit, userid, offset, name)
		if err != nil {
			return nil, err
		}
		return product, nil

	} else { //is guest
		product, err := p.productData.SelectByName(limit, offset, sortStr, name)
		if err != nil {
			return nil, err
		}
		return product, nil
	}
}

func (p *productService) UpdateById(id uint, userid uint, input product.Core, file io.Reader, handlerFilename string) (string, error) {
	if id <= 0 {
		return "", errors.New(errid)
	}

	if file != nil && handlerFilename != "" {
		timestamp := time.Now().Unix()
		fileName := fmt.Sprintf("%d_%s", timestamp, handlerFilename)
		photoFileName, errPhoto := p.helper.UploadProductPicture(file, fileName)
		if errPhoto != nil {
			return "", errPhoto
		}
		input.ProductPicture = photoFileName
	}

	err := p.productData.PutById(id, userid, input)
	if err != nil {
		return "", err
	}
	return input.ProductPicture, nil
}

func (p *productService) Delete(id uint, userid uint) error {
	if id <= 0 {
		return errors.New(errid)
	}
	return p.productData.Delete(id, userid)
}
