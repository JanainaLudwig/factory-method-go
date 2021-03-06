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
		return NewBoletoPayment(purchase), nil
	case Pix:
		return NewPixPayment(purchase), nil
	default:
		return nil, errors.New("[PAYMENT ERROR] Payment method not implemented")
	}
}
