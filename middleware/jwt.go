package middleware

import (
	"e-detect/model"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
	"time"
)

type jwtMiddleware struct {
	key string
}

type JWTService interface {
	GenerateToken(model.User) (string, error)
	UserJwtMiddleware() echo.MiddlewareFunc
	AdminJwtMiddleware() echo.MiddlewareFunc
}

// enkapsulasi jwt
func NewJwtService(secretKey string) JWTService {
	return &jwtMiddleware{
		key: secretKey,
	}
}

// fungsi middleware echo untuk dipasang pada route
//func (h *jwtMiddleware) JwtMiddleware() echo.MiddlewareFunc {
//	return middleware.JWTWithConfig(middleware.JWTConfig{
//		SigningMethod: "HS256",
//		SigningKey:    []byte(h.key),
//	})
//}

func (h *jwtMiddleware) UserJwtMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			signature := strings.Split(c.Request().Header.Get("Authorization"), " ")
			if len(signature) < 2 {
				return c.JSON(http.StatusForbidden, "Invalid token")
			}

			if signature[0] != "Bearer" {
				return c.JSON(http.StatusForbidden, "Invalid token")
			}

			claim := jwt.MapClaims{}

			token, _ := jwt.ParseWithClaims(signature[1], claim, func(t *jwt.Token) (interface{}, error) {
				return []byte(h.key), nil
			})

			method, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok || method != jwt.SigningMethodHS256 {
				return c.JSON(http.StatusForbidden, "Invalid token")
			}

			c.Set("userID", fmt.Sprintf("%v", claim["id"]))

			return next(c)
		}
	}
}

func (h *jwtMiddleware) AdminJwtMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			signature := strings.Split(c.Request().Header.Get("Authorization"), " ")
			if len(signature) < 2 {
				return c.JSON(http.StatusForbidden, "Invalid token")
			}

			if signature[0] != "Bearer" {
				return c.JSON(http.StatusForbidden, "Invalid token")
			}

			claim := jwt.MapClaims{}

			token, _ := jwt.ParseWithClaims(signature[1], claim, func(t *jwt.Token) (interface{}, error) {
				return []byte(h.key), nil
			})

			if claim["role"] == "user" {
				return echo.ErrForbidden
			}

			method, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok || method != jwt.SigningMethodHS256 {
				return c.JSON(http.StatusForbidden, "Invalid token")
			}

			c.Set("payload", fmt.Sprintf("%s", claim["id"]))

			return next(c)
		}
	}
}

// fungsi generate token ketika login user
func (h *jwtMiddleware) GenerateToken(data model.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claims["id"] = data.ID
	claims["role"] = data.Role
	key := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return key.SignedString([]byte(h.key))
}
