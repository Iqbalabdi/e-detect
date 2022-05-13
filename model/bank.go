package model

import "gorm.io/gorm"

type Bank struct {
	*gorm.Model
	Name   string
	Code   int
	Report []Report `gorm:"foreignKey:BankID"`
}
