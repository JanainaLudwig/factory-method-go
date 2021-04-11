package payment

import "errors"

type PaymentEnum int

const (
	Boleto PaymentEnum = iota
	Pix
	CreditCard
	PicPay
)

type Purchase struct {
	ValueInCents int
	BuyerEmail   string
	Method       PaymentEnum
}

type Payment interface {
	Pay() error
}

func PaymentFactory(purchase Purchase) (Payment, error) {
	switch purchase.Method {
	case Boleto:
		return &BoletoPayment{BuyInfo: purchase}, nil
	case Pix:
		return &PixPayment{BuyInfo: purchase}, nil
	default:
		return nil, errors.New("[PAYMENT ERROR] Invalid payment method")
	}
}
