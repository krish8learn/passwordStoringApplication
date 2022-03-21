package dto

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/krish8learn/passwordStoringApplication/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	err error
)

func DbConnect() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Info,      // Log level
			IgnoreRecordNotFoundError: false,            // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,             // Disable color
		},
	)
	// dsn := "host=localhost dbname=password_store_application user=root password=krish@knight8 port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(config.LodaEnvVariables()), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalln("error while DB connection", err)
	}
	fmt.Printf("DB connectiont to Successfull\n")
}
