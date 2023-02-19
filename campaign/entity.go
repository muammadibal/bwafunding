package campaign

import "time"

type Campaign struct {
	ID               int
	UserId           int32
	Name             string `gorm:"type:varchar(255)"`
	ShortDescription string `gorm:"type:varchar(255)"`
	Description      string `gorm:"type:varchar(255)"`
	Perks            string `gorm:"type:varchar(255)"`
	BackerAmount     int32
	GoalAmount       int32
	CurrentAmount    int32
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type CampaignImage struct {
	ID         int
	CampaignId int32
	FileName   string `gorm:"type:varchar(255)"`
	IsPrimary  int8
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
