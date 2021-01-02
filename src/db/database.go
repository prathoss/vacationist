package db

import (
	"fmt"

	"github.com/Prathoss/vacationist/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func Connect() {
	db, err := gorm.Open(postgres.Open(config.GetDbConnection()), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Could not connect to the database: %e", err))
	}

	db.AutoMigrate(&User{}, &Vacation{})
	Db = db
}
