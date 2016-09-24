package nisas

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/marcalegal/api/apimux"
)

// Service ...
func Service(db *gorm.DB) apimux.Service {
	return apimux.Service{
		{
			Name:        "Index",
			Method:      "GET",
			Path:        "/{term}",
			HandlerFunc: handler(db),
		},
	}
}

func handler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		var nisas []string
		term := "%" + mux.Vars(r)["term"] + "%"

		rows, _ := db.
			Table("words").
			Where("word LIKE ?", term).
			Select("word").
			Rows()
		defer rows.Close()
		for rows.Next() {
			var word string
			rows.Scan(&word)
			nisas = append(nisas, word)
		}

		response, err := json.Marshal(nisas)
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
