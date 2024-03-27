package router

import (
	"database/sql"
	"sayakaya-test/pkg/controller"
	"sayakaya-test/pkg/repository"
	"sayakaya-test/pkg/service"

	"github.com/labstack/echo/v4"
)

func NewPromoRouter(e *echo.Echo, g *echo.Group, db *sql.DB) {

	pr := repository.NewPromoRepository(db)
	ps := service.NewPromoService(pr)
	pc := &controller.PromoController{
		PromoService: ps,
	}

	g.GET("/promo", pc.CheckPromoCode)
}
