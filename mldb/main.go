package mldb

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
)

// Builddb ...
func Builddb() {
	dbURL := os.Getenv("DATABASE_URL")
	fmt.Println(dbURL)
	if dbURL == "" {
		host := "localhost"
		user := "RodrigoFuenzalida"
		dbname := "marcalegal"
		sslmode := "disable"

		dbURL = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s", host, user, dbname, sslmode)
	}

	db, err := gorm.Open("postgres", dbURL)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	db.LogMode(true)
	db.DropTableIfExists(&User{})
	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
		db.Model(&User{}).AddIndex("idx_user_email", "email")
	}

	db.DropTableIfExists(&Brand{})
	if !db.HasTable(&Brand{}) {
		db.CreateTable(&Brand{})
	}

	db.DropTableIfExists(&Natural{})
	if !db.HasTable(&Natural{}) {
		db.CreateTable(&Natural{})
	}

	db.DropTableIfExists(&Juridica{})
	if !db.HasTable(&Juridica{}) {
		db.CreateTable(&Juridica{})
	}

	db.DropTableIfExists(&Values{})
	if !db.HasTable(&Values{}) {
		db.CreateTable(&Values{})
	}
	err = CreateValues(db)
	if err != nil {
		panic(err)
	}

	db.DropTableIfExists(&Words{})
	if !db.HasTable(&Words{}) {
		db.CreateTable(&Words{})
	}

	db.DropTableIfExists(&Nisas{})
	if !db.HasTable(&Nisas{}) {
		db.CreateTable(&Nisas{})
	}

	db.DropTableIfExists(&Classes{})
	if !db.HasTable(&Classes{}) {
		db.CreateTable(&Classes{})
	}
}

// CreateValues ...
func CreateValues(db *gorm.DB) (err error) {
	tx := db.Begin()
	// Note the use of tx as the database handle once you are within a transaction

	if err := tx.Create(&Values{
		Kind:     "domainCl",
		Amount:   "12.000",
		RateKind: "pesos",
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&Values{
		Kind:     "domainCom",
		Amount:   "15.000",
		RateKind: "pesos",
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&Values{
		Kind:     "char",
		Amount:   "19",
		RateKind: "pesos",
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&Values{
		Kind:     "UF",
		Amount:   "26.224,30",
		RateKind: "pesos",
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&Values{
		Kind:     "UTM",
		Amount:   "45.999,00",
		RateKind: "pesos",
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&Values{
		Kind:     "costo_publi",
		Amount:   "2.636",
		RateKind: "pesos",
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&Values{
		Kind:     "costo_mix",
		Amount:   "10.643",
		RateKind: "pesos",
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&Values{
		Kind:     "honorarios",
		Amount:   "4",
		RateKind: "UF",
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&Values{
		Kind:     "imp_ini",
		Amount:   "1",
		RateKind: "UTM",
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&Values{
		Kind:     "imp_fin",
		Amount:   "2",
		RateKind: "UTM",
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&Values{
		Kind:     "round",
		Amount:   "100",
		RateKind: "pesos",
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
