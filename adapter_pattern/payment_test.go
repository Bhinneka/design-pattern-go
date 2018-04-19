package adapter

import (
  "testing"
)

func TestAdapter(t *testing.T) {

	bankTransfer := BankTransfer{}

	t.Run("Positive test", func(t *testing.T) {

		input1 := "BCA VA"
		input2 := "MANDIRI VA"
		input3 := "BNI TRANSFER"
		input4 := "BRI TRANSFER"

		expectedResult1 := "Pay With BCA VA"
		expectedResult2 := "Pay With MANDIRI VA"
		expectedResult3 := "Pay With BNI TRANSFER"
		expectedResult4 := "Pay With BRI TRANSFER"

		result1 := bankTransfer.Pay(input1)
		result2 := bankTransfer.Pay(input2)
		result3 := bankTransfer.Pay(input3)
		result4 := bankTransfer.Pay(input4)

		if result1 != expectedResult1 {
			t.Error("Should return BCA VA Payment")
		}

		if result2 != expectedResult2 {
			t.Error("Should return MANDIRI VA Payment")
		}

		if result3 != expectedResult3 {
			t.Error("Should return BNI TRANSFER Payment")
		}

		if result4 != expectedResult4 {
			t.Error("Should return BRI TRANSFER Payment")
		}
	})
}
