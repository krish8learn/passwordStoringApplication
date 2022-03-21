package dto

import (
	"fmt"
	"log"

	"github.com/krish8learn/passwordStoringApplication/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DbConnect() {
	// dsn := "host=localhost dbname=password_store_application user=root password=krish@knight8 port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(config.LodaEnvVariables()), &gorm.Config{})
	if err != nil {
		log.Fatalln("error while DB connection", err)
	}
	fmt.Printf("DB connectiont to Successfull\n")
}
