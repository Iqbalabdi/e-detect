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

type DisclaimerHandler struct {
	DUseCase model.DisclaimerUseCase
	DJwt     middleware.JWTService
}

func NewDisclaimerHandler(uc model.DisclaimerUseCase, jwt middleware.JWTService) *DisclaimerHandler {
	return &DisclaimerHandler{
		DUseCase: uc,
		DJwt:     jwt,
	}
}

func (r *DisclaimerHandler) Route(e *echo.Echo) {

	e.POST("/akun/sanggahan/buat", r.SaveDisclaimer, r.DJwt.UserJwtMiddleware())
	e.GET("/akun/sanggahan/riwayat", r.GetDisclaimerHistoryByUser, r.DJwt.UserJwtMiddleware())
	e.PUT("/akun/sangghan/edit/:id", r.UpdateDisclaimerByID)
	e.DELETE("/akun/sanggahan/:id", r.DeleteDisclaimerByID)
	e.POST("/admin/sanggahan/validasi/:id", r.DisclaimerValidating)
	e.GET("/admin/sanggahan/all", r.GetAllReport, r.DJwt.AdminJwtMiddleware())
}

func (r *DisclaimerHandler) SaveDisclaimer(c echo.Context) (err error) {
	var disclaimer model.Disclaimer
	err = c.Bind(&disclaimer)

	userID := c.Get("userID")
	disclaimer.UserID, err = strconv.Atoi(userID.(string))

	err = r.DUseCase.SaveRequest(disclaimer)

	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, disclaimer)
}

func (r *DisclaimerHandler) GetDisclaimerHistoryByUser(c echo.Context) error {

	userID, err := strconv.Atoi(c.Get("userID").(string))
	listDisclaimer, err := r.DUseCase.ReadUserDisclaimer(userID)
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, listDisclaimer)
}

func (r *DisclaimerHandler) UpdateDisclaimerByID(c echo.Context) error {

	var disclaimer model.Disclaimer
	c.Bind(&disclaimer)

	id, _ := strconv.Atoi(c.Param("id"))

	Disclaimer, err := r.DUseCase.EditDisclaimer(id, disclaimer)
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, Disclaimer)
}

func (r *DisclaimerHandler) DeleteDisclaimerByID(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	err := r.DUseCase.DeleteDisclaimer(id)
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "delete success",
	})
}

func (r *DisclaimerHandler) DisclaimerValidating(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := r.DUseCase.Validate(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "sucess validating Disclaimers with id: " + strconv.Itoa(id),
	})
}

func (r *DisclaimerHandler) GetAllReport(c echo.Context) error {
	listDisclaimer, err := r.DUseCase.GetAllDisclaimer()
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, listDisclaimer)
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
