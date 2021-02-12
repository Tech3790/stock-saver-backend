package functions

import (
	"encoding/json"
	"net/http"

	"github.com/Tech3790/stock-saver-backend/src/models"
	"github.com/gorilla/mux"
)

// UpdateUser is a method to update a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	var user models.User
	db.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	db.Save(&user)
	json.NewEncoder(w).Encode(user)
}
