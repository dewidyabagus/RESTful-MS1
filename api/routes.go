package api

import (
	"RESTfulMS1/api/v1/user"

	echo "github.com/labstack/echo/v4"
)

func RegisterRouters(e *echo.Echo, user *user.Controller) {

	if user == nil {
		panic("Error Initiate Routes")
	}

	userGroup := e.Group("/v1/users")
	userGroup.POST("/register", user.AddNewUser)
}
