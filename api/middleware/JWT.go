package middleware

import (
	"RESTfulMS1/business"
	"RESTfulMS1/config"

	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"net/http"
)

func init() {
	middleware.ErrJWTMissing = echo.NewHTTPError(
		http.StatusBadRequest,
		echo.Map{
			"code":    "400",
			"message": "Missing or Malformed JWT",
		},
	)

	middleware.ErrJWTInvalid = echo.NewHTTPError(
		http.StatusUnauthorized,
		echo.Map{
			"code":    "401",
			"message": "Invalid or Expired JWT",
		},
	)
}

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(config.GetJWTSecretKey()),
	})
}

func ExtractJWTUserId(c echo.Context) (string, error) {
	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return "", business.ErrUnauthorized
	}

	claim := user.Claims.(jwt.MapClaims)

	id, found := claim["id"]
	if !found {
		return "", business.ErrUnauthorized
	}

	return id.(string), nil
}
