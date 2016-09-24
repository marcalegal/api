package brands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/marcalegal/api/utils"
	"github.com/marcalegal/mldb"
)

// MLRequest ....
type MLRequest struct {
	Brand mldb.Brand    `json:"brand"`
	User  mldb.Natural  `json:"user"`
	RLP   mldb.Juridica `json:"rlp"`
}

func create(db *gorm.DB) http.HandlerFunc {
	return utils.BearerAuth(db, func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("UserID")
		id, _ := strconv.Atoi(userID)
		fmt.Println(id)
		var newRequest MLRequest

		if err := json.NewDecoder(r.Body).Decode(&newRequest); err != nil {
			reportError(w, err)
			return
		}
		if newRequest.RLP.Email != "" {
			fmt.Println("Save the RLP")
		}

		if newRequest.User.Email != "" {
			fmt.Println("Save the Natural")
		}

		if newRequest.Brand.Name != "" {
			fmt.Println("Save the Queen")
		}
		fmt.Println()
		// res := db.Create(&newBrand)
		// brand := res.Value.(*mldb.Brand)
		// reportError(w, res.Error)
		//
		// fmt.Println(brand.ID)
		// fmt.Println(brand.UserID)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
	})
}
