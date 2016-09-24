package indicadores

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/marcalegal/api/apimux"
	"github.com/marcalegal/api/utils/crawler"
	"github.com/marcalegal/mldb"
)

// BASEURL ...
const (
	BASEURL = "http://si3.bcentral.cl/Indicadoressiete/secure/IndicadoresDiarios.aspx"
)

type indicadores struct {
	UF  mldb.Values `json:"UF"`
	UTM mldb.Values `json:"UTM"`
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
		var indic indicadores
		var uf mldb.Values
		var utm mldb.Values
		ufValue := crawler.UF(BASEURL)
		// UF
		db.Model(&uf).Where("id = 4").Update("amount", ufValue)
		utmValue := crawler.UTM(BASEURL)
		// UTM
		db.Model(&utm).Where("id = 5").Update("amount", utmValue)

		indic.UF = uf
		indic.UTM = utm
		response, err := json.Marshal(indic)
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
