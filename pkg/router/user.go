package router

import (
	"database/sql"
	"sayakaya-test/pkg/controller"
	"sayakaya-test/pkg/repository"
	"sayakaya-test/pkg/service"

	"github.com/labstack/echo/v4"
)

func NewUserRouter(e *echo.Echo, g *echo.Group, db *sql.DB) {

	ur := repository.NewUserRepository(db)
	uu := service.NewUserService(ur)
	uc := &controller.UserController{
		UserService: uu,
	}

	g.GET("/user", uc.FetchUsers)
}
