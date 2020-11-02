package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string
	DiscordNumber string
	PingList []Ping `gorm:"foreignKey:UserId"`
}