package data

import (
	order "PetPalApp/features/order"
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type OrderModel struct {
	db *gorm.DB
}

func New(db *gorm.DB) order.OrderModel {
	return &OrderModel{
		db: db,
	}
}

func (om *OrderModel) CreateOrder(opCore order.OrderCore) error {
    var product order.Product
    if err := om.db.First(&product, "id = ?", opCore.ProductID).Error; err != nil {
        return fmt.Errorf("failed to find product with ID %d: %v", opCore.ProductID, err)
    }

    totalPrice := product.Price * float64(opCore.Quantity)

    op := Order{
        UserID:   		opCore.UserID,
        ProductID: 		opCore.ProductID,
		ProductName: 	product.ProductName,
		ProductPicture: product.ProductPicture,
        Quantity:  		opCore.Quantity,
        Price:     		totalPrice,
		Status:         "Created",
        InvoiceID:      generateInvoiceID(),
    }

    tx := om.db.Create(&op)
    return tx.Error
}

func (om *OrderModel) GetOrdersByUserID(userID uint) ([]order.OrderCore, error) {
    var orders []Order
    tx := om.db.Where("user_id = ?", userID).Find(&orders)
    if tx.Error != nil {
        return nil, tx.Error
    }

    var result []order.OrderCore
    for _, order := range orders {
        result = append(result, ToCore(order))
    }

    return result, nil
}

func (om *OrderModel) GetProductById(productID uint) (*order.Product, error) {
	var result order.Product
	if err := om.db.Where("id = ?", productID).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil 
}

func generateInvoiceID() string {
	randomNumber := rand.Intn(9000) + 1000
	currentDate := time.Now().Format("02012006")
	invoiceID := fmt.Sprintf("ORDER-%s-%d", currentDate, randomNumber)

	return invoiceID
}