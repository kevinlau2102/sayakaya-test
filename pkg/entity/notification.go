package entity

type Notification struct {
	ID               int    `json:"id"`
	NotificationType string `json:"notification_type"`
	Subject          string `json:"subject"`
	Message          string `json:"message"`
	UserID           int    `json:"user_id"`
	CreatedAt        string `json:"created_at"`
}

type NotificationRepository interface {
	CreateNotification(notification Notification) error
}

type NotificationService interface {
	CreateNotification(user User, code string) error
}
