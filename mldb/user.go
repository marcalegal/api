package mldb

import (
	"regexp"

	"github.com/jinzhu/gorm"
)

// SecretKey ...
const SecretKey = "WOW,MuchShibe,ToDogge"

// User ...
type User struct {
	gorm.Model
	Email        string `gorm:"type:varchar(30);unique"json:"email"`
	Password     string `gorm:"type:varchar(127)"json:"password"`
	Name         string `gorm:"type:varchar(255)"json:"name"`
	Lastname     string `gorm:"type:varchar(255)"json:"lastname"`
	Address      string `gorm:"type:varchar(255)"json:"address"`
	Phone        string `gorm:"type:varchar(255)"json:"phone"`
	Kind         int    `sql:"DEFAULT:'0'"gorm:"type:int"json:"kind"`
	SessionToken string `gorm:"type:varchar(255)"json:"session_token"`
}

// UserResponse ...
type UserResponse struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	Lastname     string `json:"lastname"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	SessionToken string `json:"session_token"`
	Errors       map[string]string
}

// Validate ...
func (u *UserResponse) Validate() bool {
	u.Errors = make(map[string]string)

	if u.Name == "" {
		u.Errors["Name"] = "Name can't be blank"
	}
	if u.Lastname == "" {
		u.Errors["Lastname"] = "Lastname can't be blank"
	}
	re := regexp.MustCompile(".+@.+\\..+")
	matched := re.Match([]byte(u.Email))
	if matched == false {
		u.Errors["Email"] = "Invalid email address"
	}
	if u.Password == "" {
		u.Errors["Password"] = "Password can't be blank"
	}
	if u.Phone == "" {
		u.Errors["Phone"] = "Phone can't be blank"
	}

	return len(u.Errors) == 0
}
