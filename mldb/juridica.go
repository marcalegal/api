package mldb

import "github.com/jinzhu/gorm"

// Juridica ...
type Juridica struct {
	gorm.Model
	NameRlp string `gorm:"type:varchar(255)"json:"name_rlp"`
	RutRlp  string `gorm:"type:varchar(255)"json:"rut_rlp"`
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
	Brand   Brand
	BrandID int `gorm:"type:int REFERENCES brands(id)"json:"brand_id"`
}

// JuridicaResponse ...
type JuridicaResponse struct {
	Errors map[string]string
}
