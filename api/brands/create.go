package brands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/marcalegal/api/utils"
	"github.com/marcalegal/api/utils/payment_code"
	"github.com/marcalegal/api/utils/pdf"
	"github.com/marcalegal/mldb"
)

// MLRequest ....
type MLRequest struct {
	Brand mldb.Brand    `json:"brand"`
	User  mldb.Natural  `json:"user"`
	Doms  mldb.Domains  `json:"dominios"`
	RLP   mldb.RPL      `json:"rlp"`
	Empr  mldb.Juridica `json:"empr"`
}

// CreateResponse ...
type CreateResponse struct {
	PaymentCode string `json:"reservationCode"`
	PDFURL      string `json:"url"`
	Kind        string `json:"kind"`
	BrandID     uint   `json:"brand_id"`
}

func create(db *gorm.DB) http.HandlerFunc {
	return utils.BearerAuth(db, func(w http.ResponseWriter, r *http.Request) {
		userID := r.Header.Get("UserID")
		id, _ := strconv.Atoi(userID)

		var newRequest MLRequest
		var response CreateResponse
		var currentUser mldb.User

		if err := json.NewDecoder(r.Body).Decode(&newRequest); err != nil {
			reportError(w, err)
			return
		}

		if err := db.Where("id = ?", id).First(&currentUser).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if newRequest.Brand.Name != "" {
			newRequest.Brand.UserID = uint(id)

			if err := db.Create(&newRequest.Brand).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Println("Queen is safe(save)")
			response.PaymentCode = paymentcode.Gen()
			db.
				Model(&newRequest.Brand).
				Update("payment_code", response.PaymentCode)
		}

		if newRequest.Brand.DomRegister {
			newRequest.Doms.BrandID = newRequest.Brand.ID
			if err := db.Create(&newRequest.Doms).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		fmt.Println("newRequest.RLP")
		fmt.Println(newRequest.RLP)

		if newRequest.Empr.Email != "" {
			newRequest.Empr.BrandID = newRequest.Brand.ID
			if err := db.Create(&newRequest.Empr).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if newRequest.RLP.Email != "" {
				newRequest.RLP.JuridicaID = newRequest.Empr.ID
				fmt.Println(newRequest.RLP)
				if err := db.Create(&newRequest.RLP).Error; err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				fmt.Println("RLP Saved")
			}

			url, err := pdf.Legal(
				newRequest.Brand.Name,
				newRequest.Empr,
				newRequest.RLP,
				id,
				newRequest.Brand,
			)
			response.PDFURL = url
			response.Kind = "Juridica"
			response.BrandID = newRequest.Brand.ID
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if err := db.Model(&newRequest.Brand).Updates(map[string]interface{}{
				"attorney_power": url,
				"register_kind":  "juridica",
			}).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			fmt.Println("Juridica Saved")
		}

		if newRequest.User.Email != "" {
			url, err := pdf.Natural(
				newRequest.Brand.Name,
				newRequest.User,
				id,
				newRequest.Brand,
			)
			response.PDFURL = url
			response.Kind = "Natural"
			response.BrandID = newRequest.Brand.ID
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			db.Model(&newRequest.Brand).Updates(map[string]interface{}{
				"attorney_power": url,
				"register_kind":  "natural",
			})

			newRequest.User.BrandID = newRequest.Brand.ID
			if err := db.Create(&newRequest.User).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Println("Natural Saved")
		}

		// should return PaymentCode, pdf url and
		res, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	})
}
