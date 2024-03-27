package repository

import (
	"database/sql"
	"sayakaya-test/pkg/entity"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) FetchUsers() ([]entity.User, error) {
	sql := "SELECT id, name, email, verified_status, birthday, is_birthday, created_at FROM users WHERE verified_status = true"
	rows, err := ur.db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.VerifiedStatus, &user.Birthday, &user.IsBirthday, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepository) UpdateIsBirthday(id int, isBirthday bool) error {
	sql := "UPDATE users SET is_birthday = $1 WHERE id = $2"
	stmt, err := ur.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(isBirthday, id)
	if err != nil {
		return err
	}


	return nil
}
