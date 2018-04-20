package adapter

import (
	"fmt"
	"strings"
)

//BankTransfer
type BankTransfer struct {
}

//Pay
func (p *BankTransfer) Pay(name string) string {
	var result string
	if strings.Contains(name, "TRANSFER") {
		result = fmt.Sprintf("Pay With %s", name)
	} else if name == "MANDIRI VA" || name == "BCA VA" {
		paymentAdapter := NewPaymentAdapter(name)
		result = paymentAdapter.Pay(name)
	}
	return result
}
