package register

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/marcalegal/api/utils"
	"github.com/marcalegal/api/utils/emails"
	"github.com/marcalegal/mldb"
)

func create(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		var newUser mldb.UserResponse

		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			text := fmt.Sprintf(`{
				"error": "%s"
			}`, err.Error())
			response := []byte(text)
			w.Write(response)
			return
		}

		if !newUser.Validate() {
			w.WriteHeader(http.StatusBadRequest)
			response, _ := json.Marshal(newUser.Errors)
			w.Write(response)
			return
		}

		dbUser := &mldb.User{
			Name:         newUser.Name,
			Lastname:     newUser.Lastname,
			Email:        newUser.Email,
			Password:     newUser.Password,
			Phone:        newUser.Phone,
			SessionToken: "",
		}

		fmt.Println()
		userExists := db.Where(mldb.User{
			Email: dbUser.Email,
		}).Select("email").Find(&dbUser).RowsAffected

		if userExists == 0 {
			db.Create(&dbUser)
			tokenString, err := utils.CreateToken(dbUser.ID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			var fullname = dbUser.Name + " " + dbUser.Lastname
			var email = dbUser.Email
			var password = dbUser.Password

			if !emails.RegisterEmail(fullname, email, password) {
				fmt.Println("Error sending registration emails")
			}

			db.Model(&dbUser).Update("session_token", tokenString)
			w.WriteHeader(http.StatusCreated)
			response, _ := json.Marshal(dbUser)
			w.Write(response)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		response := []byte(`{
			"error": "User already exists"
		}`)
		w.Write(response)
	}
}
