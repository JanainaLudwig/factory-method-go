package main

import (
	"golang-factory/payment"
	"log"
)

func main() {
	purchases := []payment.Purchase{
		{
			ValueInCents: 5000,
			BuyerEmail:   "email1@example.com",
			Method:       payment.Boleto,
		},
		{
			ValueInCents: 72000,
			BuyerEmail:   "email2@example.com",
			Method:       payment.Pix,
		},
		{
			ValueInCents: 12000,
			BuyerEmail:   "email3@example.com",
			Method:       payment.PicPay,
		},
	}

	for _, purchase := range purchases {
		payment, err := payment.PaymentFactory(purchase)

		if err != nil {
			log.Println(err)
			continue
		}

		payErr := payment.Pay()

		if payErr != nil {
			log.Println(err)
		}
	}
}
