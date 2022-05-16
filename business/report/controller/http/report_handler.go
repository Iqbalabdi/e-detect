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

	e.POST("/akun/laporan/rekening", r.SaveBankReport, r.RJwt.UserJwtMiddleware())
	e.POST("/akun/laporan/telepon", r.SavePhoneReport, r.RJwt.UserJwtMiddleware())
	e.GET("/akun/laporan/riwayat", r.GetReportHistoryByUser, r.RJwt.UserJwtMiddleware())
	e.PUT("/akun/laporan/:id", r.UpdateReportByID)
	e.DELETE("/akun/laporan/:id", r.DeleteReportByID)
	e.GET("/cek/statistik", r.Statistic, r.RJwt.UserJwtMiddleware())
	e.GET("/cek/rekening/:number", r.detectBank)
	e.GET("/cek/phone/:number", r.detectPhone)
	e.POST("/admin/laporan/validasi/:id", r.ReportValidating)
	e.GET("/admin/laporan/all", r.GetAllReport, r.RJwt.AdminJwtMiddleware())
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

	userID := c.Get("userID")
	report.UserID, err = strconv.Atoi(userID.(string))
	report.TipeLaporan = "rekening"

	err = r.RUseCase.SaveRequest(report)
	fmt.Println("anjay")
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, report)
}

func (r *ReportHandler) SavePhoneReport(c echo.Context) (err error) {
	var report model.Report
	err = c.Bind(&report)

	userID := c.Get("userID")
	report.UserID, err = strconv.Atoi(userID.(string))
	report.TipeLaporan = "phone"

	err = r.RUseCase.SaveRequest(report)
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, report)
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

func (r *ReportHandler) Statistic(c echo.Context) error {
	totalReport, totalBank, totalPhone, totalCost, err := r.RUseCase.Statistic()
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"totalReport": totalReport,
		"totalBank":   totalBank,
		"totalPhone":  totalPhone,
		"totalCost":   totalCost,
	})
}

func (r *ReportHandler) detectBank(c echo.Context) error {
	number := c.Param("number")
	bank, err := r.RUseCase.DetectBank(number)
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"terlapor": bank,
	})
	//return func {
	//
	//}

}

func (r *ReportHandler) detectPhone(c echo.Context) error {
	number := c.Param("number")
	phone, err := r.RUseCase.DetectPhone(number)
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"terlapor": phone,
	})
}

func (r *ReportHandler) ReportValidating(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := r.RUseCase.Validate(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "sucess validating reports with id: " + strconv.Itoa(id),
	})
}

func (r *ReportHandler) GetAllReport(c echo.Context) error {
	listReport, err := r.RUseCase.GetAllReport()
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, listReport)
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
