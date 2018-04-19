package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {

	t.Run("Positive test", func(t *testing.T) {
		instance := GetInstance()

		expectedResult := "Message : from Bhinneka"

		if instance.ShowMessage() != expectedResult {
			t.Error("Positive test error")
		}
	})
}
