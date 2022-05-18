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
	e.PUT("/akun/sangghan/edit/:id", r.UpdateDisclaimerByID, r.DJwt.UserJwtMiddleware())
	e.DELETE("/akun/sanggahan/:id", r.DeleteDisclaimerByID, r.DJwt.UserJwtMiddleware())
	e.POST("/admin/sanggahan/validasi/:id", r.DisclaimerValidating, r.DJwt.AdminJwtMiddleware())
	e.GET("/admin/sanggahan/all", r.GetAllDsclaimer, r.DJwt.AdminJwtMiddleware())
}

// SaveDisclaimer godoc
// @Summary      Save Report
// @Description  Create report
// @Tags         disclaimers
// @Accept       json
// @Produce      json
// @Success      200	{object}	response.ApiResponse
// @Failure      404	{object}	response.ApiResponse
// @Failure      500	{object}	response.ApiResponse
// @Router       /akun/sanggahan/buat [post]
func (r *DisclaimerHandler) SaveDisclaimer(c echo.Context) (err error) {
	var disclaimer model.Disclaimer
	err = c.Bind(&disclaimer)

	userID := c.Get("userID")
	disclaimer.UserID, err = strconv.Atoi(userID.(string))

	err = r.DUseCase.SaveRequest(disclaimer)

	if err != nil {
		return c.JSON(GetStatusCode(err), response.ApiResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, response.ApiResponse{
		Message: "Success",
		Data:    disclaimer,
	})
}

// GetDisclaimerHistoryByUser godoc
// @Summary      Get Disclaimer History
// @Description  Retrive all users history of disclaimer
// @Tags         disclaimers
// @Accept       json
// @Produce      json
// @Success      200	{object}	response.ApiResponse
// @Failure      404	{object}	response.ApiResponse
// @Failure      500	{object}	response.ApiResponse
// @Router       /akun/sanggahan/riwayat [get]
func (r *DisclaimerHandler) GetDisclaimerHistoryByUser(c echo.Context) error {

	userID, err := strconv.Atoi(c.Get("userID").(string))
	listDisclaimer, err := r.DUseCase.ReadUserDisclaimer(userID)
	if err != nil {
		return c.JSON(GetStatusCode(err), response.ApiResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.ApiResponse{
		Message: "Success",
		Data:    listDisclaimer,
	})
}

// UpdateDisclaimerByID godoc
// @Summary      Update disclaimer
// @Description  Edit user disclaimer by disclaimer_id
// @Tags         disclaimers
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "id"
// @Success      200	{object}	response.ApiResponse
// @Failure      404	{object}	response.ApiResponse
// @Failure      500	{object}	response.ApiResponse
// @Router       /akun/sangghan/edit/:id [put]
func (r *DisclaimerHandler) UpdateDisclaimerByID(c echo.Context) error {

	var disclaimer model.Disclaimer
	c.Bind(&disclaimer)

	id, _ := strconv.Atoi(c.Param("id"))

	Disclaimer, err := r.DUseCase.EditDisclaimer(id, disclaimer)
	if err != nil {
		return c.JSON(GetStatusCode(err), response.ApiResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.ApiResponse{
		Message: "Success",
		Data:    Disclaimer,
	})
}

// DeleteDisclaimerByID godoc
// @Summary      Delete disclaimer
// @Description  Delete discliamer by diclaimer_id
// @Tags         disclaimers
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "id"
// @Success      200	{object}	response.ApiResponse
// @Failure      404	{object}	response.ApiResponse
// @Failure      500	{object}	response.ApiResponse
// @Router       /akun/sanggahan/:id [delete]
func (r *DisclaimerHandler) DeleteDisclaimerByID(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	err := r.DUseCase.DeleteDisclaimer(id)
	if err != nil {
		return c.JSON(GetStatusCode(err), response.ApiResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.ApiResponse{
		Message: "Success",
	})
}

// DisclaimerValidating godoc
// @Summary      Validate disclaimer
// @Description  Validating disclaimer from Admin by disclaimer_id
// @Tags         disclaimers
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "id"
// @Success      200	{object}	response.ApiResponse
// @Failure      404	{object}	response.ApiResponse
// @Failure      500	{object}	response.ApiResponse
// @Router       /admin/sanggahan/validasi/:id [post]
func (r *DisclaimerHandler) DisclaimerValidating(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := r.DUseCase.Validate(id)
	if err != nil {
		return c.JSON(GetStatusCode(err), response.ApiResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.ApiResponse{
		Message: "success validating Disclaimers with id: " + strconv.Itoa(id),
	})
}

// GetAllDsclaimer godoc
// @Summary      Get All Disclaimer
// @Description  Get all disclaimer from admin
// @Tags         disclaimers
// @Accept       json
// @Produce      json
// @Success      200	{object}	response.ApiResponse
// @Failure      404	{object}	response.ApiResponse
// @Failure      500	{object}	response.ApiResponse
// @Router       /admin/sanggahan/all [get]
func (r *DisclaimerHandler) GetAllDsclaimer(c echo.Context) error {
	listDisclaimer, err := r.DUseCase.GetAllDisclaimer()
	if err != nil {
		return c.JSON(GetStatusCode(err), response.ApiResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.ApiResponse{
		Message: "success",
		Data:    listDisclaimer,
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
