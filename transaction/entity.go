package transaction

import "time"

type Transaction struct {
	ID         int
	CampaignId int32
	UserId     int32
	Amount     int32
	Status     string `gorm:"type:varchar(255)"`
	Code       string `gorm:"type:varchar(255)"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
