package register

import (
	"net/http"

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
			Path:        "/",
			HandlerFunc: create(db),
		},
		{
			Name:        "Remove",
			Method:      "DELETE",
			Path:        "/{userID}",
			HandlerFunc: delete(db),
		},
	}
}

func delete(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userID := vars["userID"]

		db.Where("id = ?", userID).Unscoped().Delete(mldb.User{})
	}
}
