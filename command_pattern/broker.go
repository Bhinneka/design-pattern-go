package command

//Broker
type Broker struct {
	Orders []Order
}

//AddOrder
func (b *Broker) AddOrder(order Order) {
	b.Orders = append(b.Orders, order)
}

//PlaceOrders
func (b *Broker) PlaceOrders() {
	for _, v := range b.Orders {
		v.Execute()
	}
}
