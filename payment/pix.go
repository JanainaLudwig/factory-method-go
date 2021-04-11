package payment

import "log"

type PixPayment struct {
	BuyInfo Purchase
}

func (up *PixPayment) Pay() error {
	log.Println("Paying by pix |", "User: "+up.BuyInfo.BuyerEmail, "Value: R$", (up.BuyInfo.ValueInCents / 100))

	return nil
}
