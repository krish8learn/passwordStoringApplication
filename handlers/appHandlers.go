package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krish8learn/passwordStoringApplication/dto"
)

type AppRequest struct {
	// AppID         int    `json:"app_id"`
	AppName       string `json:"app_name"`
	Reason        string `json:"reason"`
	EmailIDUsed   string `json:"email_id_used"`
	BrowserStored string `json:"browser_stored"`
	Password      string `json:"password"`
	// CreatedAt     time.Time      `json:"created_at"`
}

func SaveApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var app AppRequest

	//check whether inputs are allright
	decodeError := json.NewDecoder(r.Body).Decode(&app)
	if decodeError != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	arg := dto.Applications{
		AppName: app.AppName,
		Reason: sql.NullString{
			String: app.Reason,
			Valid:  true,
		},
		EmailIDUsed:   app.EmailIDUsed,
		BrowserStored: app.BrowserStored,
		Password:      app.Password,
	}

	DbData, DbErr := dto.CreateApp(arg)
	if DbErr != nil {
		// fmt.Println(DbErr.Error())
		if DbErr.Error() == `ERROR: duplicate key value violates unique constraint "emails_pkey" (SQLSTATE 23505)` {
			respondWithError(w, http.StatusConflict, "Database Issue")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Database Issue")
		return
	}
	json.NewEncoder(w).Encode(DbData)
	return

}

func GetAppDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	DbApp, DbEr := dto.GetApp(params["id"])
	if DbEr != nil {
		if DbEr == sql.ErrNoRows {
			respondWithError(w, http.StatusNotFound, "No data found")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Database Issue")
		return
	}

	respondWithJSON(w, http.StatusOK, DbApp)
}

func GetAllApps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	DbApps, DbErr := dto.GetAllApps()
	if DbErr != nil {
		respondWithError(w, http.StatusInternalServerError, "Database Issue")
		return
	}
	json.NewEncoder(w).Encode(DbApps)
}

func RemoveApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	DbEr := dto.RemoveApp(params["id"])
	if DbEr != nil {
		if DbEr.Error() == "No Record found" {
			respondWithError(w, http.StatusNotFound, "Not Record found")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Database Issue")
		return
	}

	respondWithJSON(w, http.StatusOK, "Deletion Successfull")
}
