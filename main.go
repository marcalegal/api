package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/marcalegal/api"
	"github.com/marcalegal/mldb"
)

func main() {
	if create := os.Getenv("CREATE_DB"); create == "create" {
		mldb.Builddb()
	}

	fmt.Println("db built")
	app := api.App()
	port := os.Getenv("PORT")
	if os.Getenv("PORT") == "" {
		port = "8000"
	}
	port = fmt.Sprintf(":%s", port)
	fmt.Println("running service on port ", port)
	http.ListenAndServe(port, app)
}
