package adapter

import (
	"fmt"
)

//Payment
type Payment interface {
	Pay(string) string
}

//VirtualAccountPayment
type VirtualAccount interface {
	PayWithBCAVA() string
	PayWithMandiriVA() string
}

//BCAVirtualAcount
type BCAVirtualAccount struct {
	Name string
}

//PayWithBCAVA
func (v *BCAVirtualAccount) PayWithBCAVA() string {
	return fmt.Sprintf("Pay With %s", v.Name)
}

//PayWithMandiriVA
func (v *BCAVirtualAccount) PayWithMandiriVA() string {
	return ""
}

//MandiriVirtualAccount
type MandiriVirtualAccount struct {
	Name string
}

//PayWithBCAVA
func (v *MandiriVirtualAccount) PayWithBCAVA() string {
	return ""
}

//PayWithMandiriVA
func (v *MandiriVirtualAccount) PayWithMandiriVA() string {
	return fmt.Sprintf("Pay With %s", v.Name)
}
