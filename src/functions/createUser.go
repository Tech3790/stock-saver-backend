package functions

import (
	"encoding/json"
	"net/http"

	"github.com/Tech3790/stock-saver-backend/src/models"
)

// CreateUser is a method to create a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)
	json.NewEncoder(w).Encode(user)
}
