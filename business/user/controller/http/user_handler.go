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

func (u *UserHandler) GetAll(c echo.Context) error {
	listUs, err := u.UUsecase.GetAll()
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, listUs)
}

func (u *UserHandler) Create(c echo.Context) (err error) {
	var newUs model.User
	err = c.Bind(&newUs)

	user, err := u.UUsecase.Create(newUs)
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, user)
}

// handler login
func (u *UserHandler) Login(c echo.Context) error {
	var userLogin model.LoginRequest
	var data model.User
	err := c.Bind(&userLogin)
	var val bool

	data, val, err = u.UUsecase.Login(userLogin)
	if err != nil {
		return c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
	}
	if val == false {
		return c.JSON(http.StatusUnauthorized, ResponseError{Message: "Unauthorized"})
	}
	token, e := u.UJwt.GenerateToken(data)
	if e != nil {
		fmt.Println("masuk error", e)
		return c.JSON(GetStatusCode(e), ResponseError{Message: e.Error()})
	}
	return c.JSON(GetStatusCode(err), ResponseError{Message: token})
}

func (u *UserHandler) Update(c echo.Context) (err error) {
	// your solution here
	var user model.User

	err = c.Bind(&user)
	if err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))

	user, err = u.UUsecase.Update(id, user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user with id : " + strconv.Itoa(id),
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
