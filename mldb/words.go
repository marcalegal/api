package mldb

// Words ...
type Words struct {
	ID   int    `gorm:"type:int"json:"id"`
	Word string `gorm:"type:varchar(255)"json:"word"`
}

// WordsResponse ...
type WordsResponse struct {
	Errors map[string]string
}
