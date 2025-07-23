package pg

import (
	"framework/domain"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func NewDB() (*DB, error) {
	db, err := gorm.Open(postgres.Open(viper.GetString("postgres.dsn")), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// auto migrate
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&domain.Route{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&domain.RouteNode{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&domain.UserTrip{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&domain.Moment{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&domain.Comment{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&domain.UserLocation{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&domain.Follow{})
	if err != nil {
		return nil, err
	}

	return &DB{DB: db}, nil
}
