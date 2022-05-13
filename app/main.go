package main

import (
	_userHttpDelivery "e-detect/business/user/controller/http"
	_userRepo "e-detect/business/user/repository/mysql"
	_userUcase "e-detect/business/user/usecase"
	"e-detect/config"
	_jwtUsecase "e-detect/middleware"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()

	// inisiasi jwtusecase
	jwtUsecase = _jwtUsecase.NewJwtService("secret")

	userRepository = _userRepo.NewMysqlUserRepository(db)
	userUsecase    = _userUcase.NewUserUseCase(userRepository)
	userHandler    = _userHttpDelivery.NewUserHandler(userUsecase, jwtUsecase)
)

func main() {
	e := echo.New()
	userHandler.Route(e)
	e.Logger.Fatal(e.Start("localhost:9090"))
}
