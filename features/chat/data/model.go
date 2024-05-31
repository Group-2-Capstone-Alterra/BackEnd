package data

import (
	"gorm.io/gorm"
)

type Chat struct {
    gorm.Model
    SenderID   uint   
    ReceiverID uint   
    Message    string 
}