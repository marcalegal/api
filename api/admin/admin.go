package admin

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/marcalegal/api/admin/brands"
	"github.com/marcalegal/api/admin/statistics"
	"github.com/marcalegal/api/apimux"
)

// Service ...
func Service(db *gorm.DB) apimux.Service {
	return apimux.Service{
		{
			Name:        "Index",
			Method:      "GET",
			Path:        "/brands",
			HandlerFunc: brands.Index(db),
		},
		{
			Name:        "Statistics",
			Method:      "GET",
			Path:        "/statistics",
			HandlerFunc: statistics.Index(db),
		},
		// {
		// 	Name:        "Show",
		// 	Method:      "GET",
		// 	Path:        "/{brandID}",
		// 	HandlerFunc: show(db),
		// },
		// {
		// 	Name:        "Create",
		// 	Method:      "POST",
		// 	Path:        "/",
		// 	HandlerFunc: create(db),
		// },
		// {
		// 	Name:        "Update",
		// 	Method:      "PATCH",
		// 	Path:        "/{brandID}",
		// 	HandlerFunc: update(db),
		// },
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
