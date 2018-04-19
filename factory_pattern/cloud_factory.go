package factory

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
