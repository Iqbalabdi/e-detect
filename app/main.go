package main

import (
	_disclaimerHttpDelivery "e-detect/business/disclaimer/controller/http"
	_disclaimertRepo "e-detect/business/disclaimer/repository/mysql"
	_disclaimerUcase "e-detect/business/disclaimer/usecase"
	_reportHttpDelivery "e-detect/business/report/controller/http"
	_reportRepo "e-detect/business/report/repository/mysql"
	_reportUcase "e-detect/business/report/usecase"
	_userHttpDelivery "e-detect/business/user/controller/http"
	_userRepo "e-detect/business/user/repository/mysql"
	_userUcase "e-detect/business/user/usecase"
	"net/http"

	"e-detect/config"
	_jwtUsecase "e-detect/middleware"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_ "e-detect/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
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

	disclaimerRepository = _disclaimertRepo.NewMysqlDisclaimerRepository(db)
	disclaimerUsecase    = _disclaimerUcase.NewDisclaimerUseCase(disclaimerRepository)
	disclaimerHandler    = _disclaimerHttpDelivery.NewDisclaimerHandler(disclaimerUsecase, jwtUsecase)
)

// @title e-detect
// @version 1.0
// @description Aplikasi untuki mendeteksi nomor rekening dan telepon yang melakukan penipuan
// @contact.name API Support
// @contact.email support@swagger.io
// @BasePath /business
func main() {
	e := echo.New()
	userHandler.Route(e)
	reportHandler.Route(e)
	disclaimerHandler.Route(e)

	handleSwag := echoSwagger.WrapHandler

	e.GET("/swagger/*", handleSwag)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "e-detect")
	})

	e.Logger.Fatal(e.Start(":9090"))
}
