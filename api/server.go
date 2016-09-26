package api

import (
	"fmt"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/marcalegal/api/admin"
	"github.com/marcalegal/api/apimux"
	"github.com/marcalegal/api/brands"
	"github.com/marcalegal/api/classes"
	"github.com/marcalegal/api/domain"
	"github.com/marcalegal/api/indicadores"
	"github.com/marcalegal/api/login"
	"github.com/marcalegal/api/logout"
	"github.com/marcalegal/api/nisas"
	"github.com/marcalegal/api/prices"
	"github.com/marcalegal/api/register"
	"github.com/marcalegal/api/uploader"
	"github.com/rs/cors"
)

// App ...
func App() *negroni.Negroni {
	r := apimux.NewRouter()

	dbURL := os.Getenv("DATABASE_URL")

	if dbURL == "" {
		host := "localhost"
		user := "RodrigoFuenzalida"
		dbname := "marcalegal"
		sslmode := "disable"

		dbURL = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s", host, user, dbname, sslmode)
	} else {
		dbURL, _ = pq.ParseURL(dbURL)
		dbURL += " sslmode=require"
	}
	db, err := gorm.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	v1 := r.AddAPIVersion(1)
	v1.AddService("/login", login.Service(db))
	v1.AddService("/logout", logout.Service(db))
	v1.AddService("/register", register.Service(db))
	v1.AddService("/brands", brands.Service(db))
	v1.AddService("/indicadores", indicadores.Service(db))
	v1.AddService("/prices", prices.Service(db))
	v1.AddService("/nisas", nisas.Service(db))
	v1.AddService("/domain", domain.Service(db))
	v1.AddService("/classes", classes.Service(db))
	v1.AddService("/uploader", uploader.Service(db))
	v1.AddService("/admin", admin.Service(db))

	// c := cors.Default()
	c := cors.New(cors.Options{
		AllowedMethods: []string{"POST", "GET", "DELETE", "PATCH", "PUT"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})
	app := negroni.New()
	app.Use(negroni.NewRecovery())
	app.Use(negroni.NewLogger())
	app.Use(c)
	app.UseHandler(r.Multiplexer())

	return app
}
