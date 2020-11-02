package models

import "gorm.io/gorm"

type Fibonacci struct {
	gorm.Model
	Input uint64
	Output uint64
}
