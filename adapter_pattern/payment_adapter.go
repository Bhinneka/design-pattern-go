package adapter

//PaymentAdapter
type PaymentAdapter struct {
	virtualAccount VirtualAccount
}

//NewPaymentAdapter
func NewPaymentAdapter(name string) *PaymentAdapter {
	var virtualAccount VirtualAccount
	if name == "BCA VA" {
		virtualAccount = &BCAVirtualAccount{"BCA VA"}
	} else if name == "MANDIRI VA" {
		virtualAccount = &MandiriVirtualAccount{"MANDIRI VA"}
	}
	return &PaymentAdapter{virtualAccount}
}

//Pay
func (p *PaymentAdapter) Pay(name string) string {
	var result string
	if name == "BCA VA" {
		result = p.virtualAccount.PayWithBCAVA()
	} else if name == "MANDIRI VA" {
		result = p.virtualAccount.PayWithMandiriVA()
	}
	return result
}
