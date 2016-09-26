package recover

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/marcalegal/api/apimux"
	"github.com/marcalegal/api/utils/emails"
)

// Service ...
func Service(db *gorm.DB) apimux.Service {
	return apimux.Service{
		{
			Name:        "Index",
			Method:      "GET",
			Path:        "/{email}",
			HandlerFunc: recover(db),
		},
	}
}

// Users ...
type Users struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func recover(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		email := mux.Vars(r)["email"]
		var user Users
		db.
			Table("users").
			Where("email = ?", email).
			Select("name, lastname, email, password").Find(&user)

		fmt.Println(user)
		fullname := fmt.Sprintf("%s %s", user.Name, user.Lastname)
		if emails.RecoverEmail(fullname, user.Email, user.Password) {
			w.WriteHeader(http.StatusOK)
			response := []byte(`{
				"message": "Email was sent"
			}`)
			w.Write(response)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		response := []byte(`{
			"error": "Could not send email."
		}`)
		w.Write(response)

	}
}
