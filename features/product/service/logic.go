package service

import (
	"PetPalApp/features/product"
	"PetPalApp/utils/helper"
	"errors"
	"fmt"
	"io"
	"log"
	"time"
)

type productService struct {
	productData product.DataInterface
	helper      helper.HelperInterface
}

func New(pd product.DataInterface, helper helper.HelperInterface) product.ServiceInterface {
	return &productService{
		productData: pd,
		helper:      helper,
	}
}

func (p *productService) Create(id uint, input product.Core, file io.Reader, handlerFilename string) (string, error) {
	if id <= 0 || input.ProductName == "" || input.Price <= 0 || file == nil {
		return "", errors.New("Nama Produk / Harga / Foto Produk tidak boleh kosong!")
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

func (p *productService) GetAll(userid uint, offset uint, sortStr string) ([]product.Core, error) {
	log.Println("[Service]")
	log.Println("[Service] sortStr", sortStr)
	product, _ := p.productData.VerIsAdmin(userid)
	if product.ID == 0 { // not admin
		log.Println("[Service - not admin]")
		product, err := p.productData.SelectAll(offset, sortStr)
		if err != nil {
			return nil, err
		}
		if userid != 0 { // add this check
			// distance
			if sortStr == "lowest distance" || sortStr == "higest distance" {
				productSort := p.helper.SortProductsByDistance(userid, product)
				log.Println("[service - not admin] distance")
				return productSort, nil
			}
		}
		log.Println("[service - not admin] not distance")
		return product, nil
	} else { //if admin
		log.Println("[Service - admin]")
		product, err := p.productData.SelectAllAdmin(userid, offset)
		if err != nil {
			return nil, err
		}
		return product, nil
	}
}

func (p *productService) GetProductById(id uint, userid uint) (data *product.Core, err error) {

	if id <= 0 {
		return nil, errors.New("[validation] product id not valid")
	}

	if userid != 0 {
		return p.productData.SelectByIdAdmin(id, userid)
	} else {
		return p.productData.SelectById(id)
	}
}

func (p *productService) UpdateById(id uint, userid uint, input product.Core, file io.Reader, handlerFilename string) (string, error) {
	if id <= 0 {
		return "", errors.New("id not valid")
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
		return errors.New("id not valid")
	}
	return p.productData.Delete(id, userid)
}
