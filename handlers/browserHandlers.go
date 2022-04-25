package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krish8learn/passwordStoringApplication/dto"
)

type BrowserRequest struct {
	BrowserName    string `json:"browser_name"`
	AccountCreated bool   `json:"account_created"`
	EmailIDUsed    string `json:"email_id_used"`
}

func SaveBrowser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var browser BrowserRequest
	decodeError := json.NewDecoder(r.Body).Decode(&browser)

	if decodeError != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}
	DbData, DbErr := dto.SaveBrowser(browser.BrowserName, browser.EmailIDUsed, browser.AccountCreated)
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

func GetAllBrowsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	DbBrowsers, DbErr := dto.GetAllBrowsers()
	if DbErr != nil {
		respondWithError(w, http.StatusInternalServerError, "Database Issue")
		return
	}
	json.NewEncoder(w).Encode(DbBrowsers)
}

func RemoveBrowser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	DbEr := dto.RemoveBrowser(params["id"])
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
