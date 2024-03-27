package repository

import (
	"database/sql"
	"sayakaya-test/pkg/entity"
	"time"
)

type NotificationRepository struct {
	db *sql.DB
}

func NewNotificationRepository(db *sql.DB) entity.NotificationRepository {
	return &NotificationRepository{
		db: db,
	}
}

func (nr *NotificationRepository) CreateNotification(req entity.Notification) error {
	sql := `INSERT INTO notifications (notification_type, subject, message, user_id, created_at) VALUES ($1, $2, $3, $4, $5)`
	stmt, err := nr.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err2 := stmt.Exec(req.NotificationType, req.Subject, req.Message, req.UserID, time.Now())
	if err2 != nil {
		return err2
	}
	return nil
}
