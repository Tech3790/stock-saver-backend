package functions

import (
	"encoding/json"
	"net/http"

	"github.com/Tech3790/stock-saver-backend/src/models"
	"github.com/gorilla/mux"
)

// DeleteUser is a method to delete a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	var user models.User
	db.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("the user is deleted!")
}
