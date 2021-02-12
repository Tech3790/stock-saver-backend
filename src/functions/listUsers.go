package functions

import (
	"encoding/json"
	"net/http"

	"github.com/Tech3790/stock-saver-backend/src/models"
)

// ListUsers is a method to list all users
func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var users []models.User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}
