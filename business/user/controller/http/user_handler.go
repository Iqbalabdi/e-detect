package http

import (
	"e-detect/business/response"
	"e-detect/middleware"
	"e-detect/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ResponseError struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UUsecase model.UserUseCase
	UJwt     middleware.JWTService
}

func NewUserHandler(uc model.UserUseCase, jwt middleware.JWTService) *UserHandler {
	return &UserHandler{
		UUsecase: uc,
		UJwt:     jwt,
	}
}

func (u *UserHandler) Route(e *echo.Echo) {
	e.GET("/users", u.GetAll, u.UJwt.AdminJwtMiddleware())
	//e.PUT("/users/:d", u.Update)
	e.POST("/akun/login", u.Login)
	e.POST("/akun/register", u.Create)
	e.PUT("/akun/edit/:id", u.Update, u.UJwt.UserJwtMiddleware())
}

// GetAll godoc
// @Summary      Get all users
// @Description  Retrieve list of all users
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200	{object}	response.ApiResponse
// @Failure      404	{object}	response.ApiResponse
// @Failure      403	{string}	string		"Unauthorized"
// @Failure      500	{object}	response.ApiResponse
// @Router       /users [get]
func (u *UserHandler) GetAll(c echo.Context) error {
	listUs, err := u.UUsecase.GetAll()
	if err != nil {
		return c.JSON(GetStatusCode(err), response.ApiResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.ApiResponse{
		Message: "success",
		Data:    listUs,
	})
}

// Create godoc
// @Summary      Create user
// @Description  create user adn save to db
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200	{object}	response.ApiResponse
// @Failure      404	{object}	response.ApiResponse
// @Failure      500	{object}	response.ApiResponse
// @Router       /akun/register [post]
func (u *UserHandler) Create(c echo.Context) (err error) {
	var newUs model.User
	err = c.Bind(&newUs)

	user, err := u.UUsecase.Create(newUs)
	if err != nil {
		return c.JSON(GetStatusCode(err), response.ApiResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.ApiResponse{
		Message: "success",
		Data:    user,
	})
}

// Login godoc
// @Summary      Login
// @Description  Login user account
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200	{object}	response.ApiResponse
// @Failure      404	{object}	response.ApiResponse
// @Failure      500	{object}	response.ApiResponse
// @Router       /akun/login [post]
func (u *UserHandler) Login(c echo.Context) error {
	var userLogin model.LoginRequest
	var data model.User
	err := c.Bind(&userLogin)
	var val bool

	data, val, err = u.UUsecase.Login(userLogin)
	if err != nil {
		return c.JSON(GetStatusCode(err), response.ApiResponse{
			Message: err.Error(),
		})
	}
	if val == false {
		return c.JSON(http.StatusUnauthorized, response.ApiResponse{
			Message: "Unauthorized",
		})
	}
	token, e := u.UJwt.GenerateToken(data)
	if e != nil {
		fmt.Println("masuk error", e)
		return c.JSON(GetStatusCode(err), response.ApiResponse{
			Message: e.Error(),
		})
	}
	return c.JSON(GetStatusCode(err), response.ApiResponse{
		Message: "success",
		Data:    token,
	})
}

// Update godoc
// @Summary      Update user
// @Description  Update user data with id parameter
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200	{string}	response.ApiResponse
// @Failure      404	{object}	response.ApiResponse
// @Failure      500	{object}	response.ApiResponse
// @Router       /akun/update/:id [put]
func (u *UserHandler) Update(c echo.Context) (err error) {
	// your solution here
	var user model.User

	err = c.Bind(&user)
	if err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))

	user, err = u.UUsecase.Update(id, user)

	return c.JSON(http.StatusOK, response.ApiResponse{
		Message: "success update user with id : " + strconv.Itoa(id),
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
