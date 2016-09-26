package emails

import (
	"fmt"
	"net/mail"
	"os"

	"github.com/hjr265/postmark.go/postmark"
	"github.com/marcalegal/mldb"
)

// SendEmailNatural ...pos
func SendEmailNatural(path string, user mldb.Natural, brand mldb.Brand) bool {
	data := map[string]interface{}{
		"finalPrice":      brand.Total,
		"reservationCode": brand.PaymentCode,
	}
	c := postmark.Client{
		ApiKey: "f7b61ac2-0b1a-4e85-93da-08890fce306b",
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
			Address: "rf@finciero.com",
		},
		To: []*mail.Address{
			{
				Name:    fullname,
				Address: user.Email,
			},
		},
		TemplateId:    929062,
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
	data := map[string]interface{}{
		"finalPrice":      brand.Total,
		"reservationCode": brand.PaymentCode,
	}
	c := postmark.Client{
		ApiKey: "f7b61ac2-0b1a-4e85-93da-08890fce306b",
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
			Address: "rf@finciero.com",
		},
		To: []*mail.Address{
			{
				Name:    user.Fullname,
				Address: user.Email,
			},
		},
		TemplateId:    929062,
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
		ApiKey: "f7b61ac2-0b1a-4e85-93da-08890fce306b",
		Secure: true,
	}

	res, err := c.Send(&postmark.Message{
		From: &mail.Address{
			Name:    "Marcalegal",
			Address: "rf@finciero.com",
		},
		To: []*mail.Address{
			{
				Name:    fullname,
				Address: email,
			},
		},
		TemplateId:    955281,
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
