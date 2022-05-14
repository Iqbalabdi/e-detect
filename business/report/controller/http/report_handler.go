package http

import (
	"e-detect/middleware"
	"e-detect/model"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ReportHandler struct {
	RUseCase model.ReportUseCase
	RJwt     middleware.JWTService
}

func NewReportHandler(uc model.ReportUseCase, jwt middleware.JWTService) *ReportHandler {
	return &ReportHandler{
		RUseCase: uc,
		RJwt:     jwt,
	}
}

func (r *ReportHandler) Route(e *echo.Echo) {
	//e.GET("/user/getReport", u.GetAll, u.UJwt.JwtMiddleware())
	//e.GET("/user/getBankReport", u.Login)
	//e.GET("/user/getPhoneReport", r.Create, u.UJwt.JwtMiddleware())
	e.POST("/akun/laporan/rekening", r.SaveBankReport)
	e.POST("/akun/laporan/telepon", r.SavePhoneReport)
	e.GET("/akun/laporan/riwayat", r.GetReportHistoryByUser, r.RJwt.UserJwtMiddleware())
	e.PUT("/akun/laporan/:id", r.UpdateReportByID)
	e.DELETE("/akun/laporan/:id", r.DeleteReportByID)
	//e.GET("/cek/statistic", r.Statistic)
	//e.GET("/user/getReportUser", r.GetReportByUser, r.RJwt.AdminJwtMiddleware(), r.RJwt.AdminJwtMiddleware())
}

func (r *ReportHandler) Save(c echo.Context) (err error) {
	userID := c.Get("id")
	var NewReport model.Report
	err = c.Bind(&NewReport)
	NewReport.UserID, err = strconv.Atoi(userID.(string))
	user, err := r.RUseCase.SaveRequest(NewReport)
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, user)
}

func (r *ReportHandler) GetReportHistoryByUser(c echo.Context) error {

	listReport, err := r.RUseCase.ReadUserReports()
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, listReport)
}

func (r *ReportHandler) SaveBankReport(c echo.Context) (err error) {
	var report model.Report
	err = c.Bind(&report)

	report.TipeLaporan = "rekening"

	//userID := c.Get("id")
	//report.UserID, err = strconv.Atoi(userID.(string))

	data, err := r.RUseCase.SaveRequest(report)
	fmt.Println("anjay")
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, data)
}

func (r *ReportHandler) SavePhoneReport(c echo.Context) error {
	var report model.Report
	err := c.Bind(&report)

	report.TipeLaporan = "phone"

	userID := c.Get("id")
	report.UserID, _ = strconv.Atoi(userID.(string))

	data, err := r.RUseCase.SaveRequest(report)
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, data)
}

func (r *ReportHandler) UpdateReportByID(c echo.Context) error {
	var report model.Report
	c.Bind(&report)

	id, _ := strconv.Atoi(c.Param("id"))

	report, err := r.RUseCase.EditReport(id, report)
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, report)
}

func (r *ReportHandler) DeleteReportByID(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	err := r.RUseCase.DeleteReport(id)
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "delete success",
	})
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case model.ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
