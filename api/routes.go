package api

import (
	"RESTfulMS1/api/v1/auth"
	"RESTfulMS1/api/v1/user"

	echo "github.com/labstack/echo/v4"
)

func RegisterRouters(e *echo.Echo, user *user.Controller, auth *auth.Controller) {

	if user == nil || auth == nil {
		panic("Error Initiate Routes")
	}

	userGroup := e.Group("/v1/users")
	userGroup.POST("/register", user.AddNewUser)
	userGroup.POST("/login", auth.LoginUserWithEmailPassword)
}
