package brands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/marcalegal/api/utils"
	"github.com/marcalegal/mldb"
)

func show(db *gorm.DB) http.HandlerFunc {
	return utils.BearerAuth(db, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		brandID := mux.Vars(r)["brandID"]
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

		var brands mldb.BrandResponse
		db.Where(
			"id = ? and user_id = ?",
			brandID,
			id,
		).Order("created_at desc").Find(&brands.Brand)

		var nat mldb.Natural
		var jur mldb.Juridica
		var rpl mldb.RPL
		var doms mldb.DomainsResponse

		if brands.Brand.RegisterKind == "natural" {
			db.
				Table("naturals").
				Where("brand_id = ?", brands.Brand.ID).
				Find(&nat)
			brands.Rut = nat.Rut
			brands.Username = fmt.Sprintf("%s %s", nat.Name, nat.Lastname)
			brands.Email = nat.Email
		} else {
			db.
				Table("juridicas").
				Select("id").
				Where("brand_id = ?", brands.Brand.ID).
				Find(&jur)

			db.Table("rpls").Where("juridica_id = ?", jur.ID).Find(&rpl)
			brands.Rut = rpl.Rut
			brands.Username = rpl.Fullname
			brands.Email = rpl.Email
		}

		if brands.Brand.DomRegister {
			db.
				Table("domains").
				Where("brand_id = ?", brands.Brand.ID).
				Find(&doms)
			brands.Doms = doms
		}

		res, _ := json.Marshal(brands)
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	})
}
