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
	if userid != 0 {
		return p.productData.SelectAllAdmin(userid, offset)
	}
	return p.productData.SelectAll(offset, sortStr)
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

func (p *productService) UpdateById(id uint, userid uint, input product.Core) error {
	if id <= 0 {
		return errors.New("id not valid")
	}
	err := p.productData.PutById(id, userid, input)
	if err != nil {
		return err
	}
	return nil
}

func (p *productService) Delete(id uint, userid uint) error {
	if id <= 0 {
		return errors.New("id not valid")
	}
	return p.productData.Delete(id, userid)
}
