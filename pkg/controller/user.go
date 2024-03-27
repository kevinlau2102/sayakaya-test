package controller

import (
	"net/http"
	"sayakaya-test/pkg/entity"
	"sayakaya-test/util"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService entity.UserService
}

func (uc *UserController) FetchUsers(c echo.Context) error {
	resp, err := uc.UserService.FetchUsers()
	if err != nil {
		return util.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.SetResponse(c, http.StatusOK, "success", resp)
}
