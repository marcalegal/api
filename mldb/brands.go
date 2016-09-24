package mldb

import "github.com/jinzhu/gorm"

// Brand ...
type Brand struct {
	gorm.Model
	UserID          int     `gorm:"type:int REFERENCES users(id)"json:"id_usuario"`
	NSolicitud      string  `sql:"DEFAULT:'XXXXXXX'"gorm:"type:varchar(20)"json:"n_solicitud"`
	RegisterKind    string  `gorm:"type:varchar(255)"json:"register_kind"`
	Name            string  `gorm:"type:varchar(50)"json:"name"`
	Kind            string  `gorm:"type:varchar(20)"json:"tipo"`
	Logo            string  `gorm:"type:varchar(255)"json:"imagen"`
	LogoDescription string  `gorm:"type:text"json:"descripcionLogo"`
	Conflict        bool    `gorm:"type:boolean"json:"conflict"`
	Resume          string  `gorm:"type:text"json:"resumen"`
	Total           float64 `gorm:"type:decimal"json:"finalprice"`
	DomRegister     bool    `gorm:"type:boolean"json:"dom_register"`
	State           int     `sql:"DEFAULT:0"gorm:"type:int"json:"state"`
	AttorneyPower   string  `gorm:"type:text"json:"attorney_power"`
	PaymentCode     string  `gorm:"type:text"json:"payment_code"`
	ClassJSON       string  `gorm:"type:string"json:"clases"`
	User            User
}

// BrandResponse ...
type BrandResponse struct {
	Errors map[string]string
}
