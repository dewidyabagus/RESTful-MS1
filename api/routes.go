package api

import (
	"RESTfulMS1/api/middleware"
	"RESTfulMS1/api/v1/auth"
	"RESTfulMS1/api/v1/user"

	echo "github.com/labstack/echo/v4"
)

func RegisterRouters(e *echo.Echo, user *user.Controller, auth *auth.Controller) {

	if user == nil || auth == nil {
		panic("Error Initiate Routes")
	}

	// Register and Login User
	e.POST("/register", user.AddNewUser)
	e.POST("/login", auth.LoginUserWithEmailPassword)

	// User Information Changes
	userGroup := e.Group("/v1/users")
	userGroup.Use(middleware.JWTMiddleware())
	userGroup.GET("", user.FindUserByUserId)
}
