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
	router.HandleFunc("/browser/saveBrowser", handlers.SaveBrowser).Methods("POST")
	router.HandleFunc("/browser/getAllBrowser", handlers.GetAllBrowsers).Methods("GET")
	router.HandleFunc("/browser/removeBrowser/{id}", handlers.RemoveBrowser).Methods("DELETE")
	//applications
	router.HandleFunc("/app/saveApp", handlers.SaveApp).Methods("POST")
	router.HandleFunc("/app/getApp/{id}", handlers.GetAppDetails).Methods("GET")
	router.HandleFunc("/app/getAllApp", handlers.GetAllApps).Methods("GET")
	router.HandleFunc("/app/removeApp/{id}", handlers.RemoveApp).Methods("DELETE")

	log.Fatalln(http.ListenAndServe(":8080", router))
}
