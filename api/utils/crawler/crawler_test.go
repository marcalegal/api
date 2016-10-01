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

func TestDomainCl(t *testing.T) {
	url := fmt.Sprintf(
		"http://nic.cl/registry/Whois.do?d=%s&buscar=Submit",
		"homy.cl",
	)

	res := DomainCl(url)

	if !res {
		t.Errorf("DomainCl should return true")
	}
}

func TestDomainCl2(t *testing.T) {
	url := fmt.Sprintf(
		"http://nic.cl/registry/Whois.do?d=%s&buscar=Submit",
		"rodrwan.cl",
	)

	res := DomainCl(url)

	if res {
		t.Errorf("DomainCl should return false")
	}
}
