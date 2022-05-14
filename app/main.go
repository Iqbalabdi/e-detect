package main

import (
	_userHttpDelivery "e-detect/business/user/controller/http"
	_userRepo "e-detect/business/user/repository/mysql"
	_userUcase "e-detect/business/user/usecase"

	_reportHttpDelivery "e-detect/business/report/controller/http"
	_reportRepo "e-detect/business/report/repository/mysql"
	_reportUcase "e-detect/business/report/usecase"

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

	reportRepository = _reportRepo.NewMysqlReportRepository(db)
	reportUsecase    = _reportUcase.NewReportUseCase(reportRepository)
	reportHandler    = _reportHttpDelivery.NewReportHandler(reportUsecase, jwtUsecase)
)

func main() {
	e := echo.New()
	userHandler.Route(e)
	reportHandler.Route(e)
	e.Logger.Fatal(e.Start("localhost:9090"))
}
