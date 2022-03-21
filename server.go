package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krish8learn/passwordStoringApplication/dto"
	"github.com/krish8learn/passwordStoringApplication/handlers"
)

func main() {
	//initialize/connect with the DB
	dto.DbConnect()
	//initializing the router
	// dao.CreateEmail("krish9857@outlook.com", "outlook", "123afaf", "mailoffice")
	router := mux.NewRouter()
	//endpoints
	log.Println("Server Running")
	router.HandleFunc("/healthStatus", handlers.HealthStatus).Methods("GET")
	//email
	router.HandleFunc("/savePassword", handlers.SavePassword).Methods("POST")
	router.HandleFunc("/getPassword/{id}", handlers.GetPassword).Methods("GET")
	router.HandleFunc("/updatePassword", handlers.UpdatePassword).Methods("PUT")
	router.HandleFunc("/removePassword", handlers.RemovePassword).Methods("DELETE")
	log.Fatalln(http.ListenAndServe(":8080", router))
}
