package emails

import (
	"fmt"
	"net/mail"
	"os"

	"github.com/hjr265/postmark.go/postmark"
	"github.com/leekchan/accounting"
	"github.com/marcalegal/mldb"
)

// SendEmailNatural ...pos
func SendEmailNatural(path string, user mldb.Natural, brand mldb.Brand) bool {
	ac := accounting.Accounting{Symbol: "$", Precision: 2, Thousand: ".", Decimal: ","}
	data := map[string]interface{}{
		"finalPrice":      ac.FormatMoney(brand.Total),
		"reservationCode": brand.PaymentCode,
	}
	c := postmark.Client{
		ApiKey: "c6db85e0-809d-4cfd-a0db-1f970a714958",
		Secure: true,
	}
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	var attch = make([]postmark.Attachment, 0)
	pdf := &postmark.Attachment{
		Name:        fmt.Sprintf("poder_notarial_%s.pdf", brand.Name),
		Content:     file,
		ContentType: "application/pdf",
	}
	attch = append(attch, *pdf)

	fullname := fmt.Sprintf("%s %s", user.Name, user.Lastname)
	res, err := c.Send(&postmark.Message{
		From: &mail.Address{
			Name:    "Marcalegal",
			Address: "contacto@marcalegal.cl",
		},
		To: []*mail.Address{
			{
				Name:    fullname,
				Address: user.Email,
			},
		},
		TemplateId:    1080816,
		TemplateModel: data,
		Attachments:   attch,
	})

	if err != nil {
		panic(err)
	}

	if res.ErrorCode == 0 {
		fmt.Printf("%s\n", res.Message)
		return true
	}

	return false
}

// SendEmailLegal ...pos
func SendEmailLegal(path string, user mldb.RPL, brand mldb.Brand) bool {
	ac := accounting.Accounting{Symbol: "$", Precision: 2, Thousand: ".", Decimal: ","}
	data := map[string]interface{}{
		"finalPrice":      ac.FormatMoney(brand.Total),
		"reservationCode": brand.PaymentCode,
	}
	c := postmark.Client{
		ApiKey: "c6db85e0-809d-4cfd-a0db-1f970a714958",
		Secure: true,
	}
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	var attch = make([]postmark.Attachment, 0)
	pdf := &postmark.Attachment{
		Name:        fmt.Sprintf("poder_notarial_%s.pdf", brand.Name),
		Content:     file,
		ContentType: "application/pdf",
	}
	attch = append(attch, *pdf)

	res, err := c.Send(&postmark.Message{
		From: &mail.Address{
			Name:    "Marcalegal",
			Address: "contacto@marcalegal.cl",
		},
		To: []*mail.Address{
			{
				Name:    user.Fullname,
				Address: user.Email,
			},
		},
		TemplateId:    1080816,
		TemplateModel: data,
		Attachments:   attch,
	})

	if err != nil {
		panic(err)
	}

	if res.ErrorCode == 0 {
		fmt.Printf("%s\n", res.Message)
		return true
	}

	return false
}

// RecoverEmail ///
func RecoverEmail(fullname, email, password string) bool {
	data := map[string]interface{}{
		"password": password,
		"fullname": fullname,
	}
	c := postmark.Client{
		ApiKey: "c6db85e0-809d-4cfd-a0db-1f970a714958",
		Secure: true,
	}

	res, err := c.Send(&postmark.Message{
		From: &mail.Address{
			Name:    "Marcalegal",
			Address: "contacto@marcalegal.cl",
		},
		To: []*mail.Address{
			{
				Name:    fullname,
				Address: email,
			},
		},
		TemplateId:    1080817,
		TemplateModel: data,
	})

	if err != nil {
		panic(err)
	}

	if res.ErrorCode == 0 {
		fmt.Printf("%s\n", res.Message)
		return true
	}

	return false
}

// RegisterEmail ...Register email
func RegisterEmail(fullname, email, password string) bool {
	data := map[string]interface{}{
		"name":     fullname,
		"username": email,
		"password": password,
	}
	c := postmark.Client{
		ApiKey: "c6db85e0-809d-4cfd-a0db-1f970a714958",
		Secure: true,
	}

	res, err := c.Send(&postmark.Message{
		From: &mail.Address{
			Name:    "Marcalegal",
			Address: "contacto@marcalegal.cl",
		},
		To: []*mail.Address{
			{
				Name:    fullname,
				Address: email,
			},
		},
		TemplateId:    1083300,
		TemplateModel: data,
	})

	if err != nil {
		panic(err)
	}

	if res.ErrorCode == 0 {
		fmt.Printf("%s\n", res.Message)
		return true
	}

	return false
}
