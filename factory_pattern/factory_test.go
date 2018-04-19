package factory

import (
	"testing"
)

func TestFactory(t *testing.T) {

	c := CloudFactory{}

	t.Run("Positive test", func(t *testing.T) {

		input1 := "AWS"
		input2 := "AZURE"
		input3 := "GOOGLE_CLOUD"

		expectedResult1 := "Deploy using AWS"
		expectedResult2 := "Deploy using AZURE"
		expectedResult3 := "Deploy using GOOGLE_CLOUD"

		cloud1 := c.GetCloud(input1)
		cloud2 := c.GetCloud(input2)
		cloud3 := c.GetCloud(input3)

		if cloud1.Deploy() != expectedResult1 {
			t.Error("Test 1 error")
		}

		if cloud2.Deploy() != expectedResult2 {
			t.Error("Test 2 error")
		}

		if cloud3.Deploy() != expectedResult3 {
			t.Error("Test 3 error")
		}
	})
}
