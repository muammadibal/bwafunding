package transaction

import (
	"bwafunding/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignID int32
	UserID     int32
	Amount     int32
	Status     string `gorm:"type:varchar(255)"`
	Code       string `gorm:"type:varchar(255)"`
	User       user.User
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
