package utils

import (
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/marcalegal/mldb"
)

// Info ...
type Info struct {
	Valid  bool
	UserID float64
}

const expireOffset = 3600
const adminKind = 1

func getTokenRemainingValidity(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds() + expireOffset)
		}
	}
	return expireOffset
}

// ValidateToken ...
func ValidateToken(tokenString string) *Info {
	var newInfo Info
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(mldb.SecretKey), nil
	})

	if err == nil && token.Valid {
		newInfo.Valid = true
		newInfo.UserID = token.Claims["userid"].(float64)
		return &newInfo
	}

	newInfo.Valid = false
	return &newInfo
}

// CreateToken ....
func CreateToken(userID uint) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims["userid"] = userID
	// Expire in 24 hours
	token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString([]byte(mldb.SecretKey))
	return tokenString, err
}

// BearerAuth ....
func BearerAuth(db *gorm.DB, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userToken := r.Header.Get("Authorization")
		re := regexp.MustCompile("Bearer (.+)?")
		matched := re.FindStringSubmatch(userToken)
		bearer := matched[1]
		info := ValidateToken(bearer)

		if !info.Valid {
			var user mldb.User
			db.Model(&user).
				Where("id = ?", int(info.UserID)).
				Update("session_token", "")
			http.Error(w, "Authorization failed", http.StatusUnauthorized)
			return
		}

		r.Header.Add("UserID", strconv.Itoa(int(info.UserID)))

		h(w, r)
	}
}

// AdminAuth ....
func AdminAuth(db *gorm.DB, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userToken := r.Header.Get("Authorization")
		re := regexp.MustCompile("Bearer (.+)?")
		matched := re.FindStringSubmatch(userToken)
		bearer := matched[1]
		valid := validateAdminToken(db, bearer)

		if !valid {
			http.Error(w, "Authorization failed", http.StatusUnauthorized)
			return
		}

		h(w, r)
	}
}

func validateAdminToken(db *gorm.DB, bearer string) bool {
	row := db.
		Table("users").
		Select("kind").
		Where("session_token = ?", bearer).
		Row()

	var kind int
	row.Scan(&kind)

	if kind == 1 {
		return true
	}
	return false
}
