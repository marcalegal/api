package login

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/marcalegal/api/apimux"
	"github.com/marcalegal/api/utils"
	"github.com/marcalegal/mldb"
)

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Service ...
func Service(db *gorm.DB) apimux.Service {
	return apimux.Service{
		{
			Name:        "Index",
			Method:      "POST",
			Path:        "/",
			HandlerFunc: handler(db),
		},
	}
}

func handler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		var newLogin login
		var user mldb.User
		err := json.NewDecoder(r.Body).Decode(&newLogin)
		reportError(w, err)

		db.Where("email = ? and password = ?", newLogin.Username, newLogin.Password).First(&user)

		if user.SessionToken != "" {
			w.WriteHeader(http.StatusAccepted)
			response, _ := json.Marshal(user)
			w.Write(response)
			return
		}

		if user.ID == 0 {
			w.WriteHeader(http.StatusBadRequest)
			response := `{"message": "Data is not valid"}`
			w.Write([]byte(response))
			return
		}

		tokenString, err := utils.CreateToken(user.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// str := base64.StdEncoding.EncodeToString([]byte(tokenString))
		db.Model(&user).Update("session_token", tokenString)

		response, err := json.Marshal(user)
		reportError(w, err)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
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
