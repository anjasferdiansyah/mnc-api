package domains

import "time"

type (
	Transaction struct {
		ID uint `gorm:"primary_key"`
		Title string `gorm:"title"`
		Ammount int `gorm:"ammount"`
		Description string `gorm:"description"`
		Date time.Time `gorm:"date"`
	}
)