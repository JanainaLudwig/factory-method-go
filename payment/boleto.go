package payment

import "log"

type BoletoPayment struct {
	BuyInfo Purchase
}

func (up *BoletoPayment) Pay() error {
	log.Println("Paying by boleto |", "User: "+up.BuyInfo.BuyerEmail, "Value: R$", (up.BuyInfo.ValueInCents / 100))

	return nil
}
