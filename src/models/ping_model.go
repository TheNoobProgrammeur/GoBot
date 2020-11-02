package models

import "gorm.io/gorm"

type Ping struct {
	gorm.Model
	UserId         uint
}
