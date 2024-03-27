package controller

import (
	"net/http"
	"sayakaya-test/pkg/entity"
	"sayakaya-test/util"

	"github.com/labstack/echo/v4"
)

type PromoController struct {
	PromoService entity.PromoService
}

func (pc *PromoController) CheckPromoCode(c echo.Context) error {
	promoCode := c.QueryParam("code")
	resp, err := pc.PromoService.CheckPromoCode(promoCode)
	if err != nil {
		return util.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.SetResponse(c, http.StatusOK, "success", resp)
}
