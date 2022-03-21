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
	// s := router.Host("").Subrouter()
	//endpoints
	log.Println("Server Running")
	router.HandleFunc("/healthStatus", handlers.HealthStatus).Methods("GET")
	//email
	router.HandleFunc("/email/savePassword", handlers.SavePassword).Methods("POST")
	router.HandleFunc("/email/getPassword/{id}", handlers.GetPassword).Methods("GET")
	router.HandleFunc("/email/getAllPassword", handlers.GetAllPassword).Methods("GET")
	router.HandleFunc("/email/updatePassword", handlers.UpdatePassword).Methods("PUT")
	router.HandleFunc("/email/removePassword/{id}", handlers.RemovePassword).Methods("DELETE")
	//browser
	log.Fatalln(http.ListenAndServe(":8080", router))
}
