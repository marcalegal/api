package mldb

// Nisas ...
type Nisas struct {
	ID          int    `gorm:"type:int"json:"id"`
	Kind        string `gorm:"type:varchar(255)"json:"kind"`
	NisaID      string `gorm:"type:varchar(255)"json:"nisa_id"`
	NisaSubID   string `gorm:"type:varchar(255)"json:"nisa_sub_id"`
	Description string `gorm:"type:varchar(255)"json:"description"`
}

// NisasResponse ...
type NisasResponse struct {
	Errors map[string]string
}
