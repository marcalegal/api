package mldb

import "github.com/jinzhu/gorm"

// Values ...
type Values struct {
	gorm.Model
	Kind     string `gorm:"type:varchar(255)"json:"kind"`
	Amount   string `gorm:"type:varchar(255)"json:"amount"`
	RateKind string `gorm:"type:varchar(255)"json:"rate_kind"`
}

// ValuesResponse ...
type ValuesResponse struct {
	Errors map[string]string
}
