package mldb

import "github.com/jinzhu/gorm"

// Natural ...
type Natural struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255)"json:"nombre"`
	Lastname string `gorm:"type:varchar(255)"json:"apellido"`
	Rut      string `gorm:"type:varchar(255)"json:"rut"`
	Address  string `gorm:"type:varchar(255)"json:"direccion"`
	Comuna   string `gorm:"type:varchar(255)"json:"comuna"`
	City     string `gorm:"type:varchar(255)"json:"region"`
	Country  string `gorm:"type:varchar(255)"json:"pais"`
	Email    string `gorm:"type:varchar(255)"json:"email"`
	Phone    string `gorm:"type:varchar(255)"json:"telefono"`
	BrandID  uint   `gorm:"type:int REFERENCES brands(id)"json:"brand_id"`
}

// NaturalResponse ...
type NaturalResponse struct {
	Errors map[string]string
}
