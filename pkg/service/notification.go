package service

import (
	"net/smtp"
	"sayakaya-test/config"
	"sayakaya-test/pkg/entity"
	"strconv"
)

type NotificationService struct {
	NotificationRepository entity.NotificationRepository
}

func NewNotificationService(nr entity.NotificationRepository) entity.NotificationService {
	return &NotificationService{
		NotificationRepository: nr,
	}
}

func (ns *NotificationService) CreateNotification(user entity.User, code string) error {
	notification := entity.Notification{
		NotificationType: "email",
		Subject:          "SayaKaya Birthday Promo!",
		Message:          "Happy Birthday, " + user.Name + "! Enjoy 20% discount on your next purchase! Use code: " + code,
		UserID:           user.ID,
	}

	cfg, err := config.LoadConfig(".env")
	if err != nil {
		panic(err)
	}

	auth := smtp.PlainAuth("", cfg.SMTP.User, cfg.SMTP.Pass, cfg.SMTP.Host)

	msg := "From: SayaKaya\r\n" +
		"To: " + user.Email + "\r\n" +
		"Subject: " + notification.Subject + "\r\n" +
		"\r\n" + notification.Message + "\r\n"

	err = smtp.SendMail(
		cfg.SMTP.Host+":"+strconv.Itoa(cfg.SMTP.Port),
		auth,
		cfg.SMTP.User,
		[]string{user.Email},
		[]byte(msg),
	)

	if err != nil {
		panic(err)
	}

	return ns.NotificationRepository.CreateNotification(notification)
}
