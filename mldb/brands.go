package mldb

import "github.com/jinzhu/gorm"

// Brand ...
type Brand struct {
	gorm.Model
	UserID          uint    `gorm:"type:int REFERENCES users(id)"json:"id_usuario"`
	NSolicitud      string  `sql:"DEFAULT:'XXXXXXX'"gorm:"type:varchar(20)"json:"n_solicitud"`
	RegisterKind    string  `gorm:"type:varchar(255)"json:"user_tipo"`
	Name            string  `gorm:"type:varchar(50)"json:"name"`
	Kind            string  `gorm:"type:varchar(20)"json:"tipo"`
	Logo            string  `gorm:"type:text"json:"imagen"`
	LogoDescription string  `gorm:"type:text"json:"descripcionLogo"`
	Conflict        bool    `gorm:"type:boolean"json:"conflict"`
	Resume          string  `gorm:"type:text"json:"resumen"`
	Total           float64 `gorm:"type:decimal"json:"finalprice"`
	DomRegister     bool    `gorm:"type:boolean"json:"registrodom"`
	State           int     `sql:"DEFAULT:0"gorm:"type:int"json:"state"`
	AttorneyPower   string  `gorm:"type:text"json:"attorney_power"`
	PaymentCode     string  `gorm:"type:text"json:"payment_code"`
	ClassJSON       string  `gorm:"type:text"json:"clases"`
}

// BrandErrorResponse ...
type BrandErrorResponse struct {
	Errors map[string]string
}

// BrandResponse ...
type BrandResponse struct {
	Brand    Brand           `json:"marca"`
	Doms     DomainsResponse `json:"dominios"`
	Rut      string          `json:"rut"`
	Username string          `json:"username"`
	Email    string          `json:"email"`
}
