package functions

import (
	"encoding/json"
	"net/http"

	"github.com/Tech3790/stock-saver-backend/src/models"
	"github.com/gorilla/mux"
)

// GetUser is a method to get a user
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	var user models.User
	db.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}
