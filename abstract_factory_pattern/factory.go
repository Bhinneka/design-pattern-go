package factory

//Cloud interface
type Cloud interface {
	Deploy() string
}

//CITool interface
type CITool interface {
	Build() string
}

//AbstractFactory
type AbstractFactory interface {
	GetCloud(string) Cloud
	GetCITool(string) CITool
}

//CloudFactory
type CloudFactory struct {
}

//GetCloud
func (c CloudFactory) GetCloud(name string) Cloud {
	switch name {
	case "AWS":
		return NewAWS(name)
	case "AZURE":
		return NewAzure(name)
	case "GOOGLE_CLOUD":
		return NewGoogleCloud(name)
	default:
		return nil
	}
}

//GetCITool
func (c CloudFactory) GetCITool(name string) CITool {
	return nil
}

//CIToolFactory
type CIToolFactory struct {
}

//GetCloud
func (c CIToolFactory) GetCloud(name string) Cloud {
	return nil
}

//GetCITool
func (c CIToolFactory) GetCITool(name string) CITool {
	switch name {
	case "TRAVIS":
		return NewTravis(name)
	case "JENKINS":
		return NewJenkins(name)
	case "CIRCLE":
		return NewCircle(name)
	default:
		return nil
	}
}

//FactoryGenerator
type FactoryGenerator struct {
}

//GetFactory return AbstractFactory
func (f FactoryGenerator) GetFactory(choice string) AbstractFactory {
	switch choice {
	case "CLOUD":
		return CloudFactory{}
	case "CI":
		return CIToolFactory{}
	default:
		return nil
	}
}
