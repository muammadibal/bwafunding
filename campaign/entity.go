package campaign

import (
	"bwafunding/user"
	"time"
)

type Campaign struct {
	ID               int
	UserID           int32
	Name             string `gorm:"type:varchar(255)"`
	ShortDescription string `gorm:"type:varchar(255)"`
	Description      string
	Perks            string `gorm:"type:varchar(255)"`
	BackerAmount     int32
	GoalAmount       int32
	CurrentAmount    int32
	Slug             string `gorm:"type:varchar(255)"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	User             user.User
	CampaignImages   []CampaignImage
}

type CampaignImage struct {
	ID         int
	CampaignId int32
	FileName   string `gorm:"type:varchar(255)"`
	IsPrimary  int8
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
