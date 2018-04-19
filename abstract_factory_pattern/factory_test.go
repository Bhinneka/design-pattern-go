package factory

import (
	"testing"
)

func TestFactory(t *testing.T) {

	factoryGenerator := FactoryGenerator{}

	t.Run("Cloud test", func(t *testing.T) {

		cloud := factoryGenerator.GetFactory("CLOUD")

		input1 := "AWS"
		input2 := "AZURE"
		input3 := "GOOGLE_CLOUD"

		expectedResult1 := "Deploy using AWS"
		expectedResult2 := "Deploy using AZURE"
		expectedResult3 := "Deploy using GOOGLE_CLOUD"

		cloud1 := cloud.GetCloud(input1)
		cloud2 := cloud.GetCloud(input2)
		cloud3 := cloud.GetCloud(input3)

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

	t.Run("CI Tool test", func(t *testing.T) {

		ci := factoryGenerator.GetFactory("CI")

		input1 := "TRAVIS"
		input2 := "JENKINS"
		input3 := "CIRCLE"

		expectedResult1 := "Build using TRAVIS"
		expectedResult2 := "Build using JENKINS"
		expectedResult3 := "Build using CIRCLE"

		ci1 := ci.GetCITool(input1)
		ci2 := ci.GetCITool(input2)
		ci3 := ci.GetCITool(input3)

		if ci1.Build() != expectedResult1 {
			t.Error("Test 1 error")
		}

		if ci2.Build() != expectedResult2 {
			t.Error("Test 2 error")
		}

		if ci3.Build() != expectedResult3 {
			t.Error("Test 3 error")
		}
	})
}
