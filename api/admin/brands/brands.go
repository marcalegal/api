package brands

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/marcalegal/api/utils"
	"github.com/marcalegal/mldb"
)

// Index ...
func Index(db *gorm.DB) http.HandlerFunc {
	return utils.AdminAuth(db, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		userID := r.Header.Get("UserID")
		id, err := strconv.Atoi(userID)
		if err != nil {
			// here should report error to sentry or some system like that.
			response := []byte(`{
					"error": "Cannot decode user id"
				}`)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(response)
			return
		}
		var brands []mldb.Brand
		result := db.Where("user_id = ?", id).Order("created_at desc").Find(&brands)

		res, _ := json.Marshal(result.Value)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	})
}
