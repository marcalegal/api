package logout

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/marcalegal/api/apimux"
	"github.com/marcalegal/mldb"
)

// Service ...
func Service(db *gorm.DB) apimux.Service {
	return apimux.Service{
		{
			Name:        "Index",
			Method:      "POST",
			Path:        "/{userID}",
			HandlerFunc: handler(db),
		},
	}
}

func handler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		var user mldb.User
		userID := mux.Vars(r)["userID"]

		id, err := strconv.Atoi(userID)
		reportError(w, err)
		db.Where("id = ?", id).First(&user)

		if user.SessionToken != "" {
			db.Model(&user).Update("session_token", "")
			response, _ := json.Marshal(user)
			w.WriteHeader(http.StatusOK)
			w.Write(response)
			return
		}
	}
}

func reportError(w http.ResponseWriter, err error) {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		text := fmt.Sprintf(`{
			"error": "%s"
		}`, err.Error())
		response := []byte(text)
		w.Write(response)
		return
	}
}
