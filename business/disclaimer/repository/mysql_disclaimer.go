package repository

import (
	"e-detect/model"
	"gorm.io/gorm"
)

type mysqlDisclaimerRepository struct {
	connection *gorm.DB
}

func NewMysqlDisclaimerRepository(db *gorm.DB) model.DisclaimerRepository {
	return &mysqlDisclaimerRepository{
		connection: db,
	}
}
