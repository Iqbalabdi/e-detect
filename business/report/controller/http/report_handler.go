package http

import (
	"e-detect/business/response"
	"e-detect/middleware"
	"e-detect/model"
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
	e.PUT("/akun/laporan/:id", r.UpdateReportByID, r.RJwt.UserJwtMiddleware())
	e.DELETE("/akun/laporan/:id", r.DeleteReportByID, r.RJwt.UserJwtMiddleware())
	e.GET("/cek/statistik", r.Statistic, r.RJwt.UserJwtMiddleware())
	e.GET("/cek/rekening/:number", r.DetectBank, r.RJwt.UserJwtMiddleware())
	e.GET("/cek/phone/:number", r.DetectPhone, r.RJwt.UserJwtMiddleware())
	e.PUT("/admin/laporan/validasi/:id", r.ReportValidating, r.RJwt.AdminJwtMiddleware())
	e.GET("/admin/laporan/all", r.GetAllReport, r.RJwt.AdminJwtMiddleware())
}

// GetReportHistoryByUser godoc
// @Summary      Get Report History
// @Description  Retrieve list of all users report history
// @Tags         reports
// @Accept       json
// @Produce      json
// @Success      200	{object}	response.ApiResponse
// @Failure      404	{object}	response.ApiResponse
// @Failure      500	{object}	response.ApiResponse
// @Router       /akun/laporan/riwayat [get]
func (r *ReportHandler) GetReportHistoryByUser(c echo.Context) (err error) {

	userID := c.Get("userID")
	idUser, _ := strconv.Atoi(userID.(string))

	listReport, err := r.RUseCase.ReadUserReports(idUser)
	if err != nil {
		return c.JSON(GetStatusCode(err), response.ApiResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.ApiResponse{
		Message: "Success",
		Data:    listReport,
	})
}

// SaveBankReport godoc
// @Summary      Save Bank Report
// @Description  Create report for the bank account number that commits fraud
// @Tags         reports
// @Accept       json
// @Produce      json
// @Success      200	{object}	response.ApiResponse
// @Failure      404	{object}	response.ApiResponse
// @Failure      500	{object}	response.ApiResponse
// @Router       /akun/laporan/rekening [post]
func (r *ReportHandler) SaveBankReport(c echo.Context) (err error) {
	var report model.Report
	err = c.Bind(&report)

	userID := c.Get("userID")
	report.UserID, err = strconv.Atoi(userID.(string))
	report.TipeLaporan = "rekening"

	err = r.RUseCase.SaveRequest(report)
	if err != nil {
		return c.JSON(GetStatusCode(err), response.ApiResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, response.ApiResponse{
		Message: "Success",
		Data:    report,
	})
}

// SavePhoneReport godoc
// @Summary      Save Phone Report
// @Description  Create report for the phone number that commits fraud
// @Tags         reports
// @Accept       json
// @Produce      json
// @Success      200	{object}	response.ApiResponse
// @Failure      404	{object}	response.ApiResponse
// @Failure      500	{object}	response.ApiResponse
// @Router       /akun/laporan/phone [post]
func (r *ReportHandler) SavePhoneReport(c echo.Context) (err error) {
	var report model.Report
	err = c.Bind(&report)

	userID := c.Get("userID")
	report.UserID, err = strconv.Atoi(userID.(string))
	report.TipeLaporan = "phone"

	err = r.RUseCase.SaveRequest(report)
	if err != nil {
		return c.JSON(GetStatusCode(err), response.ApiResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, response.ApiResponse{
		Message: "Success",
		Data:    report,
	})
}

// UpdateReportByID godoc
// @Summary      Update report
// @Description  Edit report by id
// @Tags         reports
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200	{object}	response.ApiResponse
// @Failure      404	{object}	response.ApiResponse
// @Failure      500	{object}	response.ApiResponse
// @Router     	 /akun/laporan/:id [put]
func (r *ReportHandler) UpdateReportByID(c echo.Context) error {
	var report model.Report
	c.Bind(&report)

	id, _ := strconv.Atoi(c.Param("id"))

	report, err := r.RUseCase.EditReport(id, report)
	if err != nil {
		return c.JSON(GetStatusCode(err), response.ApiResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.ApiResponse{
		Message: "Success Update Reports with id " + strconv.Itoa(id),
	})
}

// DeleteReportByID godoc
// @Summary      Delete report
// @Description  Delete report by id
// @Tags         reports
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200	{object}	response.ApiResponse
// @Failure      404	{object}	response.ApiResponse
// @Failure      500	{object}	response.ApiResponse
// @Router     	 /akun/laporan/:id [delete]
func (r *ReportHandler) DeleteReportByID(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	err := r.RUseCase.DeleteReport(id)
	if err != nil {
		return c.JSON(GetStatusCode(err), response.ApiResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.ApiResponse{
		Message: "delete success",
	})
}

// Statistic godoc
// @Summary      statistics
// @Description  Show statistics about bank account and phone report
// @Tags         detect
// @Accept       json
// @Produce      json
// @Success      200	{object}	response.ApiResponse
// @Failure      404	{object}	response.ApiResponse
// @Failure      500	{object}	response.ApiResponse
// @Router     	 /cek/statistik [get]
func (r *ReportHandler) Statistic(c echo.Context) error {
	totalReport, totalBank, totalPhone, totalCost, err := r.RUseCase.Statistic()
	if err != nil {
		return c.JSON(GetStatusCode(err), response.ApiResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"totalReport": totalReport,
		"totalBank":   totalBank,
		"totalPhone":  totalPhone,
		"totalCost":   totalCost,
	})
}

// DetectBank godoc
// @Summary      Detect bank
// @Description  Detect bank account who commits fraud
// @Tags         detect
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "number"
// @Success      200	{object}	response.ApiResponse
// @Failure      404	{object}	response.ApiResponse
// @Failure      500	{object}	response.ApiResponse
// @Router     	 /cek/rekening/:number [get]
func (r *ReportHandler) DetectBank(c echo.Context) error {
	number := c.Param("number")
	bank, err := r.RUseCase.DetectBank(number)
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"terlapor": bank,
	})
}

// DetectPhone godoc
// @Summary      Detect phone
// @Description  Detect phone number who commits fraud
// @Tags         detect
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "number"
// @Success      200	{object}	response.ApiResponse
// @Failure      404	{object}	response.ApiResponse
// @Failure      500	{object}	response.ApiResponse
// @Router     	 /cek/phone/:number [get]
func (r *ReportHandler) DetectPhone(c echo.Context) error {
	number := c.Param("number")
	phone, err := r.RUseCase.DetectPhone(number)
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"terlapor": phone,
	})
}

// ReportValidating godoc
// @Summary      Validate report
// @Description  Validate user report (bank account or phone number) by report id
// @Tags         reports
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "id"
// @Success      200	{object}	response.ApiResponse
// @Failure      404	{object}	response.ApiResponse
// @Failure      500	{object}	response.ApiResponse
// @Router       /admin/laporan/validasi/:id [put]
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

// GetAllReport godoc
// @Summary      Get all report
// @Description  Admin can Get all report from all users
// @Tags         reports
// @Accept       json
// @Produce      json
// @Success      200	{object}	response.ApiResponse
// @Failure      404	{object}	response.ApiResponse
// @Failure      500	{object}	response.ApiResponse
// @Router       /admin/laporan/all [get]
func (r *ReportHandler) GetAllReport(c echo.Context) error {
	listReport, err := r.RUseCase.GetAllReport()
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, response.ApiResponse{
		Message: "success",
		Data:    listReport,
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
