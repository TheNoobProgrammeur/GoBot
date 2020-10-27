package services

import (
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)


type serviceBDD struct {
	db *gorm.DB
	err error
}

var serviceInstantiated *serviceBDD


func New() *serviceBDD {
	if serviceInstantiated == nil {
		serviceInstantiated = new(serviceBDD)
	}
	return serviceInstantiated
}

func (s *serviceBDD) ConnectionBDD()  {
	dsn := "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	s.db, s.err = gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
		)

	if s.err != nil {
		panic(errors.New("Erreur dans Connection a la BDD"))
	}
}

