package repository

import (
	"database/sql"
	"sayakaya-test/pkg/entity"
	"time"
)

type PromoRepository struct {
	db *sql.DB
}

func NewPromoRepository(db *sql.DB) entity.PromoRepository {
	return &PromoRepository{
		db: db,
	}
}

func (pr *PromoRepository) GeneratePromoCode(req entity.Promo) error {
	sql := `INSERT INTO promos (code, start_date, end_date, amount, user_id, created_at) VALUES ($1, $2, $3, $4, $5, $6)`
	stmt, err := pr.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err2 := stmt.Exec(req.Code, req.StartDate, req.EndDate, req.Amount, req.UserID, time.Now())
	if err2 != nil {
		return err2
	}
	return nil
}

func (pr *PromoRepository) GetPromo(code string) (entity.Promo, error) {
	sql := `SELECT id, code, start_date, end_date, amount, user_id, created_at FROM promos WHERE code = $1`
	row := pr.db.QueryRow(sql, code)

	var promo entity.Promo
	err := row.Scan(&promo.ID, &promo.Code, &promo.StartDate, &promo.EndDate, &promo.Amount, &promo.UserID, &promo.CreatedAt)
	if err != nil {
		return promo, err
	}

	return promo, nil
}
