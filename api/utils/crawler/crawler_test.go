package crawler

import (
	"fmt"
	"testing"
)

func TestUF(t *testing.T) {
	url := "http://si3.bcentral.cl/Indicadoressiete/secure/IndicadoresDiarios.aspx"
	fmt.Println(UF(url))
}

func TestUTM(t *testing.T) {
	url := "http://si3.bcentral.cl/Indicadoressiete/secure/IndicadoresDiarios.aspx"
	fmt.Println(UTM(url))
}
