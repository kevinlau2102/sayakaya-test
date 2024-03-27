package entity

type User struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	VerifiedStatus bool   `json:"verified_status"`
	Birthday       string `json:"birthday"`
	IsBirthday     bool   `json:"is_birthday"`
	CreatedAt      string `json:"created_at"`
}

type UserRepository interface {
	FetchUsers() ([]User, error)
	UpdateIsBirthday(id int, isBirthday bool) error
}

type UserService interface {
	FetchUsers() ([]User, error)
	UpdateIsBirthday(id int, isBirthday bool) error
}
