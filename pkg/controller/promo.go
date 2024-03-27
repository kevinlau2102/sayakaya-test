package controller

import (
	"net/http"
	"sayakaya-test/pkg/entity"
	"sayakaya-test/util"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PromoController struct {
	PromoService entity.PromoService
}

func (pc *PromoController) CheckPromoCode(c echo.Context) error {
	promoCode := c.QueryParam("code")
	userID, _ := strconv.Atoi(c.QueryParam("user_id"))
	resp, err := pc.PromoService.CheckPromoCode(promoCode, userID)
	if err != nil {
		return util.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.SetResponse(c, http.StatusOK, "success", resp)
}
