package config

import (
	"e-detect/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func SetupDatabaseConnection() *gorm.DB {
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "",
		"DB_Host":     "localhost",
		"DB_Port":     "3306",
		"DB_Name":     "e-detect",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"],
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//DB.Migrator().DropTable(&model.Disclaimer{})
	DB.AutoMigrate(&model.User{}, &model.Bank{})
	DB.AutoMigrate(&model.Report{})
	DB.AutoMigrate(&model.Disclaimer{})

	return DB
}

//func CloseDatabaseConnection(db *gorm.DB) {
//	err := db.
//	if err != nil {
//		panic("failed to close database connection")
//	}
//}
