package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LodaEnvVariables() string {
	//load env files
	err := godotenv.Load("./config/app.env")
	if err != nil {
		log.Fatalln("unable to find env files", err)
	}
	log.Println("Loading env file")
	//load environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	name := os.Getenv("DB_NAME")
	pass := os.Getenv("DB_PASS")
	// dsn := "host=localhost dbname=password_store_application user=root password=krish@knight8 port=5432 sslmode=disable"
	dbConnect := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		host, port, user, name, pass)
	// fmt.Println(dbConnect)
	return dbConnect
}
