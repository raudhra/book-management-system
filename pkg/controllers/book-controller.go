package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/raudhra/book-management-system/models"
)

var NewBooks models.Books

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
