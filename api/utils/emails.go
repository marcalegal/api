package utils

import (
	"fmt"
	"net/mail"

	"github.com/hjr265/postmark.go/postmark"
)

// SendEmail ...pos
func SendEmail() bool {
	data := map[string]interface{}{
		"finalPrice":      120000,
		"reservationCode": "XDF234",
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
				Name:    "Rodrigo Fuenzalida",
				Address: "rfuenzalidac87@gmail.com",
			},
		},
		TemplateId:    929062,
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
