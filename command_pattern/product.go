package command

//Product
type Product struct {
	ID       string
	Name     string
	Quantity int
	Price    float64
}

//NewProduct
func NewProduct(id, name string, qty int, price float64) *Product {
	return &Product{ID: id, Name: name, Quantity: qty, Price: price}
}

//AddQuantity
func (p *Product) AddQuantity() {
	p.Quantity = p.Quantity + 1
}

//SubstractQuantity
func (p *Product) SubstractQuantity() {
	p.Quantity = p.Quantity - 1
}
