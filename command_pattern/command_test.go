package command

import (
	"testing"
)

func TestCommandPatter(t *testing.T) {

	product := NewProduct("001", "IPHONE X", 10, 20000000.0)

	t.Run("Positive Test", func(t *testing.T) {

		buyProduct := NewBuyProduct(product)
		sellProduct := NewSellProduct(product)

		broker := Broker{}
		broker.AddOrder(buyProduct)
		broker.AddOrder(sellProduct)

		broker.PlaceOrders()

		if product.Quantity != 10 {
			t.Error("Quantity should be 10")
		}
	})
}
