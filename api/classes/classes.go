package classes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"strings"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/marcalegal/api/apimux"
)

// SQLClass ...
type SQLClass struct {
	Class  int    `json:"class"`
	Kind   string `json:"kind"`
	Detail string `json:"detail"`
}

// SQLResponse ...
type SQLResponse []SQLClass

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
		var sqlresp SQLResponse
		term := mux.Vars(r)["term"]
		var terms = strings.Split(term, ",")
		var rows *sql.Rows

		if len(terms) == 2 {
			sanitizeTerm1 := strings.TrimSpace(terms[0])
			term1 := "%" + sanitizeTerm1 + "%"
			term2 := "%" + sanitizeTerm1 + "%"
			term3 := "%" + sanitizeTerm1 + "%"
			sanitizeTerm2 := strings.TrimSpace(terms[1])
			term4 := "%" + sanitizeTerm2 + "%"
			term5 := "%" + sanitizeTerm2 + "%"
			term6 := "%" + sanitizeTerm2 + "%"

			rows, _ = db.
				Table("words").
				Where(
					"detail LIKE ? OR detail LIKE ? OR detail LIKE ? OR detail LIKE ? OR detail LIKE ? OR detail LIKE ?",
					term1,
					term2,
					term3,
					term4,
					term5,
					term6,
				).
				Select("word").
				Rows()
		} else {
			sanitizeTerm1 := strings.TrimSpace(terms[0])
			term1 := "% " + sanitizeTerm1 + "%"
			term2 := "%" + sanitizeTerm1 + " %"
			term3 := "%" + sanitizeTerm1 + "%"
			rows, _ = db.
				Table("classes").
				Where(
					"detail LIKE ? OR detail LIKE ? OR detail LIKE ?",
					term1,
					term2,
					term3,
				).
				Select("class_id, detail, kind").
				Rows()
		}

		defer rows.Close()

		for rows.Next() {
			var classID int
			var detail, kind string
			rows.Scan(&classID, &detail, &kind)
			sqlresp = append(sqlresp, SQLClass{
				Class:  classID,
				Kind:   kind,
				Detail: detail,
			})
		}

		type Classes []string
		resp := make(map[int]Classes)
		services := make(map[int]string)
		for _, val := range sqlresp {
			text := strings.Replace(val.Detail, "\r", "", -1)
			if _, ok := resp[val.Class]; ok {
				resp[val.Class] = append(resp[val.Class], text)
				services[val.Class] = val.Kind
			} else {
				resp[val.Class] = append(resp[val.Class], text)
				services[val.Class] = val.Kind
			}
		}
		type ObjClasses map[string]string
		newResp := make(map[string]ObjClasses)

		for key, val := range resp {
			idx := 0
			nk := strconv.Itoa(key)
			if _, ok := newResp[nk]; !ok {
				newResp[nk] = make(ObjClasses)
				for k, v := range val {
					newResp[nk][strconv.Itoa(k)] = v
					idx++
				}
				newResp[nk]["cantidad"] = strconv.Itoa(idx)
				newResp[nk]["tipo"] = services[key]
				newResp[nk]["class"] = strconv.Itoa(key)
			}
		}

		response, err := json.Marshal(newResp)
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

func oneWord(word string) {

}

func twoWord(words []string) {

}
