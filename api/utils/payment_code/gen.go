package paymentcode

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const stringlength = 7

// Gen ...
func Gen() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, stringlength)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
