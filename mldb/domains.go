package mldb

import "github.com/jinzhu/gorm"

// Domains ...
type Domains struct {
	gorm.Model
	CL      string `gorm:"type:varchar(255)"json:"cl"`
	COM     string `gorm:"type:varchar(255)"json:"com"`
	BrandID uint   `gorm:"type:int REFERENCES brands(id)"json:"brand_id"`
}

// DomainsErrorResponse ...
type DomainsErrorResponse struct {
	Errors map[string]string
}

// DomainsResponse ...
type DomainsResponse struct {
	CL  string `json:"cl"`
	COM string `json:"com"`
}
