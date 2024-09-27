package models

import "time"

type User struct {
	Id             int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name           string `json:"name"`
	RoleId         string `json:"role_id"`
	Email          string `json:"email,omitempty" gorm:"type:varchar(50);unique"`
	Password       string `json:"password" gorm:"size:255"`
	Licensed       bool   `json:"licensed" gorm:"default:false"`
	LicenseEnd     time.Time
	TelegramID string `json:"telegram_id"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
