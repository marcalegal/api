package prices

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/marcalegal/api/apimux"
)

// BASEURL ...
const (
	BASEURL = "http://si3.bcentral.cl/Indicadoressiete/secure/IndicadoresDiarios.aspx"
)

type price struct {
	Amount   string `json:"amount"`
	RateKind string `json:"rate_kind"`
}

// Service ...
func Service(db *gorm.DB) apimux.Service {
	return apimux.Service{
		{
			Name:        "Index",
			Method:      "GET",
			Path:        "/",
			HandlerFunc: handler(db),
		},
	}
}

func handler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		prices := make(map[string]*price)

		rows, _ := db.Table("values").Select("kind, amount, rate_kind").Rows()
		defer rows.Close()
		for rows.Next() {
			var kind, amount, rate string
			rows.Scan(&kind, &amount, &rate)
			prices[kind] = &price{
				Amount:   amount,
				RateKind: rate,
			}
		}

		response, err := json.Marshal(prices)
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
