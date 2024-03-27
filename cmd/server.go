package cmd

import (
	"sayakaya-test/config"
	"sayakaya-test/db"
	"sayakaya-test/pkg/router"
	"sayakaya-test/pkg/scheduler"

	"github.com/labstack/echo/v4"
)

func RunServer() {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		panic(err)
	}

	e := echo.New()
	g := e.Group("/api/v1")

	Apply(e, g, cfg)
	e.Logger.Error(e.Start(":5000"))
}

func Apply(e *echo.Echo, g *echo.Group, cfg *config.Config) {
	db := db.NewInstanceDb(cfg)
	router.NewUserRouter(e, g, db)
	router.NewPromoRouter(e, g, db)

	scheduler.RunScheduler(db)
}
