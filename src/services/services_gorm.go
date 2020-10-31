package services

import (
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"myTest/src/models"
	"os"
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


	user := os.Getenv("USER_DB")
	password := os.Getenv("PASSWORD_DB")
	dbname := os.Getenv("NAME_BD")

	dsn := "host=localhost user="+user+" password="+password+" dbname="+dbname+" port=5432"
	s.db, s.err = gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
		)

	if s.err != nil {
		panic(errors.New("Erreur dans Connection a la BDD"))
	}
}

func (s *serviceBDD) MiggrationBDD(){
	_ = s.db.AutoMigrate(&models.User{})
}

