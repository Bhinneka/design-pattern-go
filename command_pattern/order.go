package command

//Order interface
type Order interface {
	Execute()
}

//BuyProduct
type BuyProduct struct {
	product *Product
}

//NewBuyProduct
func NewBuyProduct(product *Product) *BuyProduct {
	return &BuyProduct{product}
}

//Execute
func (b *BuyProduct) Execute() {
	b.product.AddQuantity()
}

//SellProduct
type SellProduct struct {
	product *Product
}

//NewSellProduct
func NewSellProduct(product *Product) *SellProduct {
	return &SellProduct{product}
}

//Execute
func (b *SellProduct) Execute() {
	b.product.SubstractQuantity()
}
