package user

import (
	"RESTfulMS1/api/v1/user/request"
	"RESTfulMS1/business/user"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type Controller struct {
	service user.Service
}

func NewController(service user.Service) *Controller {
	return &Controller{service}
}

func (c *Controller) AddNewUser(ctx echo.Context) error {
	user := new(request.RequestUser)

	var err error

	err = ctx.Bind(user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "bad request 1"})
	}

	err = c.service.AddNewUser(user.ToUserAddSpec())
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "bad request 2"})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"message": "success"})
}
