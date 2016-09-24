package uploader

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/marcalegal/api/apimux"
	"github.com/marcalegal/api/utils"
	"github.com/marcalegal/api/utils/aws"
)

// IMGResponse ...
type IMGResponse struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Service ...
func Service(db *gorm.DB) apimux.Service {
	return apimux.Service{
		{
			Name:        "Index",
			Method:      "POST",
			Path:        "/{brandName}",
			HandlerFunc: handler(db),
		},
	}
}

func handler(db *gorm.DB) http.HandlerFunc {
	s3Handler := aws.NewAWSS3Handler("marcalegal-logos")
	return utils.BearerAuth(db, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		HeaderUserID := r.Header.Get("UserID")
		brandName := mux.Vars(r)["brandName"]
		userID, _ := strconv.Atoi(HeaderUserID)

		if err := r.ParseMultipartForm(10000); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		m := r.MultipartForm

		files := m.File["file"]
		var response IMGResponse
		for i := range files {
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			//create destination file making sure the path is writeable.
			filename := files[i].Filename

			path := fmt.Sprintf("/tmp/images/%s", filename)
			dst, err := os.Create(path)
			defer dst.Close()

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//copy the uploaded file to the destination file
			if _, err = io.Copy(dst, file); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// upload image to aws s3
			URL, err := s3Handler.UploadImage(dst, filename, userID, brandName)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			response.Name = filename
			response.URL = URL

			if err := os.Remove(path); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		resp, err := json.Marshal(response)
		reportError(w, err)

		w.WriteHeader(http.StatusCreated)
		w.Write(resp)
	})
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
