package service

import (
	"PetPalApp/features/product"
	"PetPalApp/utils/helper"
	"errors"
	"fmt"
	"io"
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
	if input.IdUser <= 0 || input.ProductName == "" || input.Price <= 0 || input.ProductPicture == "" {
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
