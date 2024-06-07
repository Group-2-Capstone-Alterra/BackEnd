package data

import (
	"PetPalApp/features/order_product"
	"fmt"

	"gorm.io/gorm"
)

type OrderProductModel struct {
	db *gorm.DB
}

func New(db *gorm.DB) order_product.OrderProductModel {
	return &OrderProductModel{
		db: db,
	}
}

func (opm *OrderProductModel) CreateOrderProduct(opCore order_product.OrderProductCore) error {
    var product order_product.Product
    if err := opm.db.First(&product, "id = ?", opCore.ProductID).Error; err != nil {
        return fmt.Errorf("failed to find product with ID %d: %v", opCore.ProductID, err)
    }

    totalPrice := product.Price * float64(opCore.Quantity)

    op := OrderProduct{
        OrderID:   opCore.OrderID,
        ProductID: opCore.ProductID,
        Quantity:  opCore.Quantity,
        Price:     totalPrice, 
    }

    tx := opm.db.Create(&op)
    return tx.Error
}

func (opm *OrderProductModel) GetOrderProductsByOrderID(orderID uint) ([]order_product.OrderProductCore, error) {
	var orderProducts []OrderProduct
	tx := opm.db.Where("order_id = ?", orderID).Find(&orderProducts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var result []order_product.OrderProductCore
	for _, op := range orderProducts {
		result = append(result, ToCore(op))
	}

	return result, nil
}

func (opm *OrderProductModel) GetPriceByProductID(productID uint) (*order_product.Product, error) {
	var result order_product.Product
	if err := opm.db.Where("id = ?", productID).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil 
}
