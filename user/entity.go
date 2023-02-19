package user

import "time"

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name" gorm:"type:varchar(255)"`
	Occupation     string    `json:"occupation" gorm:"type:varchar(255)"`
	Email          string    `json:"email" gorm:"type:varchar(255)"`
	PasswordHash   string    `json:"password" gorm:"type:varchar(255)"`
	AvatarFileName string    `json:"avatar_file_name" gorm:"type:varchar(255)"`
	Role           string    `json:"role" gorm:"type:varchar(255)"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
