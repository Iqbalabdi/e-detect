package http

import (
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
	//e.GET("/user/getReport", u.GetAll, u.UJwt.JwtMiddleware())
	//e.GET("/user/getBankReport", u.Login)
	//e.GET("/user/getPhoneReport", r.Create, u.UJwt.JwtMiddleware())
	e.POST("/user/saveReport", r.Save, r.RJwt.AdminJwtMiddleware(), r.RJwt.UserJwtMiddleware())
	e.GET("/user/getReportUser", r.GetReportByUser, r.RJwt.AdminJwtMiddleware(), r.RJwt.AdminJwtMiddleware())
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

func (r *ReportHandler) GetReportByUser(c echo.Context) error {

	listUs, err := r.RUseCase.ReadUserReports()
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, listUs)
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
