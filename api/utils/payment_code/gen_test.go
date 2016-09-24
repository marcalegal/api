package paymentcode

import (
	"fmt"
	"testing"
)

func TestGen(t *testing.T) {
	s := Gen()
	fmt.Println(s)
}
