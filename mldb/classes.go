package mldb

// Classes ...
type Classes struct {
	ID         int    `gorm:"type:int"json:"id"`
	Kind       string `gorm:"type:varchar(50)"json:"kind"`
	ClassID    int    `gorm:"type:int"json:"class_id"`
	SubClassID string `gorm:"type:varchar(25)"json:"sub_class_id"`
	Detail     string `gorm:"type:text"json:"detail"`
}

// ClassesResponse ...
type ClassesResponse struct {
	Errors map[string]string
}
