package domain

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/marcalegal/api/apimux"
	"github.com/marcalegal/api/utils/crawler"
)

// BASEURL ...
const (
	BASEURL = "http://si3.bcentral.cl/Indicadoressiete/secure/IndicadoresDiarios.aspx"
)

// DomResponse ...
type DomResponse struct {
	Status    string `json:"status"`
	Domain    string `json:"domain"`
	Available bool   `json:"available"`
}

// Service ...
func Service(db *gorm.DB) apimux.Service {
	return apimux.Service{
		{
			Name:        "Index",
			Method:      "GET",
			Path:        "/{domain}/{ext}",
			HandlerFunc: handler(db),
		},
	}
}

func handler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		params := mux.Vars(r)

		if params["ext"] == "com" {
			var response DomResponse
			url := "http://ping.eu/action.php?atype=4"
			res := crawler.DomainCom(url, params["domain"])

			if !res {
				response = DomResponse{
					Status:    "success",
					Domain:    fmt.Sprintf("www.%s.com", params["domain"]),
					Available: false,
				}
			} else {
				response = DomResponse{
					Status:    "success",
					Domain:    fmt.Sprintf("www.%s.com", params["domain"]),
					Available: true,
				}
			}
			resp, err := json.Marshal(response)
			reportError(w, err)

			w.WriteHeader(http.StatusOK)
			w.Write(resp)
			return
		}
		var response DomResponse
		url := fmt.Sprintf(
			"http://nic.cl/registry/Whois.do?d=%s&buscar=Submit",
			params["domain"],
		)
		res := crawler.DomainCl(url)
		if res {
			response = DomResponse{
				Status:    "success",
				Domain:    fmt.Sprintf("www.%s.cl", params["domain"]),
				Available: false,
			}
		} else {
			response = DomResponse{
				Status:    "success",
				Domain:    fmt.Sprintf("www.%s.cl", params["domain"]),
				Available: true,
			}
		}
		resp, err := json.Marshal(response)
		reportError(w, err)

		w.WriteHeader(http.StatusOK)
		w.Write(resp)
		return
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
