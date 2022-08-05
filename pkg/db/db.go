package db

import (
	"log"
	"os"

	"github.com/adhtanjung/trackerstache_api/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Handler struct {
	DB *gorm.DB
}

var DB Handler

func Init(url string) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	db.Logger = logger.Default.LogMode(logger.Info)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	log.Println("running migration")

	db.AutoMigrate(&models.User{})

	log.Println("migration success")
	DB = Handler{DB: db}

}
