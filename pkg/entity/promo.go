package entity

type Promo struct {
	ID        int    `json:"id"`
	Code      string `json:"Code"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Amount    int    `json:"amount"`
	UserID    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type PromoRepository interface {
	GeneratePromoCode(Promo) error
	GetPromo(code string) (Promo, error)
}

type PromoService interface {
	GeneratePromoCode(userID int) (string, error)
	CheckPromoCode(code string, userID int) (Promo, error)
}
