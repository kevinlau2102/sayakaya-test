package service

import (
	"fmt"
	"math/rand"
	"sayakaya-test/pkg/entity"
	"time"
)

type PromoService struct {
	PromoRepository entity.PromoRepository
}

func NewPromoService(pr entity.PromoRepository) *PromoService {
	return &PromoService{
		PromoRepository: pr,
	}
}

func (ps *PromoService) GeneratePromoCode(userID int) (string, error) {
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	rand.Seed(time.Now().UnixNano())

	var promoCode string

	for i := 0; i < 12; i++ {
		randomIndex := rand.Intn(len(characters))
		promoCode += string(characters[randomIndex])
	}

	promo := entity.Promo{
		Code:      promoCode,
		StartDate: time.Now().Format("2006-01-02"),
		EndDate:   time.Now().Add(24 * time.Hour).Format("2006-01-02"),
		Amount:    20000,
		UserID:    userID,
	}

	if err := ps.PromoRepository.GeneratePromoCode(promo); err != nil {
		return "", err
	}

	return promoCode, nil
}

func (ps *PromoService) CheckPromoCode(code string) (entity.Promo, error) {
	promo, err := ps.PromoRepository.GetPromo(code)

	if err != nil {
		return promo, err
	}

	if promo.ID == 0 {
		return promo, fmt.Errorf("Promo code not found")
	}

	promoEndDate, _ := time.Parse("2006-01-02T15:04:05Z", promo.EndDate)

	if time.Now().After(promoEndDate) {
		return promo, fmt.Errorf("Promo code has expired")
	}

	return promo, nil

}
