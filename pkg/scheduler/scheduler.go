package scheduler

import (
	"database/sql"
	"fmt"
	"sayakaya-test/pkg/repository"
	"sayakaya-test/pkg/service"
	"time"

	"github.com/go-co-op/gocron"
)

func RunScheduler(db *sql.DB) {
	locate, _ := time.LoadLocation("Asia/Jakarta")

	ur := repository.NewUserRepository(db)
	nr := repository.NewNotificationRepository(db)
	pr := repository.NewPromoRepository(db)

	us := service.NewUserService(ur)
	ns := service.NewNotificationService(nr)
	ps := service.NewPromoService(pr)

	s := gocron.NewScheduler(locate)

	s.Every(1).Day().At("00:00").Do(func() error {
		fmt.Println("Running scheduler")
		users, err := us.FetchUsers()
		if err != nil {
			return err
		}

		for _, user := range users {
			birthday, _ := time.Parse("2006-01-02T15:04:05Z", user.Birthday)
			if birthday.Month() == time.Now().Month() && birthday.Day() == time.Now().Day() {

				if err := us.UpdateIsBirthday(user.ID, true); err != nil {
					return err
				}

				promoCode, err := ps.GeneratePromoCode(user.ID)
				if err != nil {
					return err
				}

				err = ns.CreateNotification(user, promoCode)
				if err != nil {
					return err
				}

			} else {
				err := us.UpdateIsBirthday(user.ID, false)
				if err != nil {
					return err
				}
			}
		}
		fmt.Println("Scheduler finished")

		return nil
	})

	s.WaitForScheduleAll()
	s.StartAsync()
}
