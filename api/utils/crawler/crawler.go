package crawler

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/yhat/scrape"

	"golang.org/x/net/html"
)

const (
	// BaseURL ...
	BaseURL = "http://si3.bcentral.cl/Indicadoressiete/secure/"
)

var meses = []string{
	"Enero",
	"Febrero",
	"Marzo",
	"Abril",
	"Mayo",
	"Junio",
	"Julio",
	"Agosto",
	"Septiembre",
	"Octubre",
	"Noviembre",
	"Diciembre",
}

// Crawler ...
func Crawler(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
		return nil, err
	}

	root, err := html.Parse(resp.Body)
	if err != nil {
		if err != nil {
			fmt.Println("ERROR: Failed to parse \"" + url + "\"")
			return nil, err
		}
	}

	return root, nil
}

// Post ...
func Post(url string, data url.Values) (*html.Node, error) {
	resp, err := http.PostForm(url, data)
	if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
		return nil, err
	}

	root, err := html.Parse(resp.Body)
	if err != nil {
		if err != nil {
			fmt.Println("ERROR: Failed to parse \"" + url + "\"")
			return nil, err
		}
	}

	return root, nil
}

// UF ...
func UF(url string) string {
	root, err := Crawler(url)
	if err != nil {
		return ""
	}
	link, ok := scrape.Find(root, scrape.ById("hypLnk1_1"))
	if ok {
		ufURL := BaseURL + scrape.Attr(link, "href")
		root, err := Crawler(ufURL)
		if err != nil {
			return ""
		}
		t := time.Now().Local()
		var day string
		if (t.Day() + 1) < 10 {
			day = fmt.Sprintf("0%d", t.Day()+1)
		} else {
			day = fmt.Sprintf("%d", t.Day()+1)
		}
		month := fmt.Sprintf("%d", t.Month())
		sel, _ := strconv.Atoi(month)
		id := fmt.Sprintf("gr_ctl%s_%s", day, meses[sel-1])
		link, ok := scrape.Find(root, scrape.ById(id))
		if ok {
			return scrape.Text(link)
		}
	}
	return ""
}

// UTM ...
func UTM(url string) string {
	root, err := Crawler(url)
	if err != nil {
		return ""
	}
	link, ok := scrape.Find(root, scrape.ById("hypLnk2_8"))
	if ok {
		utmURL := BaseURL + scrape.Attr(link, "href")
		root, err := Crawler(utmURL)
		if err != nil {
			return ""
		}
		t := time.Now().Local()

		month := fmt.Sprintf("%d", t.Month())
		sel, _ := strconv.Atoi(month)
		id := fmt.Sprintf("gr_ctl28_%s", meses[sel-1])
		link, ok := scrape.Find(root, scrape.ById(id))
		if ok {
			return scrape.Text(link)
		}
	}
	return ""
}

// DomainCl ...
func DomainCl(url string) bool {
	root, err := Crawler(url)
	if err != nil {
		return false
	}
	html := scrape.Text(root)
	return strings.Contains(html, "Titular:")
}

// DomainCom ...
func DomainCom(baseURL, domain string) bool {
	data := url.Values{}
	data.Set("host", domain+".com")
	data.Set("go", "Go")
	data.Set("scode", "")

	root, err := Post(baseURL, data)
	if err != nil {
		return false
	}
	html := scrape.Text(root)
	return strings.Contains(html, "No match for")
}
