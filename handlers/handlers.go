package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/krish8learn/passwordStoringApplication/dto"
)

type CreateEmail struct {
	EmailID    string `json:"email_id"`
	DomainName string `json:"domain_name"`
	Password   string `json:"password"`
	Reason     string `json:"reason"`
	// CreatedAt  time.Time      `db:"created_at"`
}

func SavePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var email CreateEmail
	decodeError := json.NewDecoder(r.Body).Decode(&email)

	if decodeError != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}
	DbData, DbErr := dto.CreateEmail(email.EmailID, email.DomainName, email.Password, email.Reason)
	if DbErr != nil {
		respondWithError(w, http.StatusInternalServerError, "Database Issue")
		return
	}
	json.NewEncoder(w).Encode(DbData)

}

func GetPassword(w http.ResponseWriter, r *http.Request) {

}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {

}

func RemovePassword(w http.ResponseWriter, r *http.Request) {

}

func HealthStatus(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Working")
	w.Write([]byte("Gorilla!\n"))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
