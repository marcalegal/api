package mldb

import "github.com/jinzhu/gorm"

// Juridica ...
type Juridica struct {
	gorm.Model
	Name    string `gorm:"type:varchar(255)"json:"name"`
	Razon   string `gorm:"type:varchar(255)"json:"razon"`
	Rut     string `gorm:"type:varchar(255)"json:"rut"`
	Giro    string `gorm:"type:varchar(255)"json:"giro"`
	Address string `gorm:"type:varchar(255)"json:"address"`
	Comuna  string `gorm:"type:varchar(255)"json:"comuna"`
	Ciudad  string `gorm:"type:varchar(255)"json:"ciudad"`
	Country string `gorm:"type:varchar(255)"json:"country"`
	Email   string `gorm:"type:varchar(255)"json:"email"`
	Phone   string `gorm:"type:varchar(255)"json:"phone"`
	BrandID uint   `gorm:"type:int REFERENCES brands(id)"json:"brand_id"`
}

// RPL ... Representante legal
type RPL struct {
	ID         uint   `gorm:"primary_key"json:"id"`
	Fullname   string `gorm:"type:varchar(255)"json:"fullname"`
	Rut        string `gorm:"type:varchar(255)"json:"rut"`
	Email      string `gorm:"type:varchar(255)"json:"email"`
	JuridicaID uint   `gorm:"type:int REFERENCES juridicas(id)"json:"juridicas_id"`
}

// JuridicaResponse ...
type JuridicaResponse struct {
	Errors map[string]string
}
