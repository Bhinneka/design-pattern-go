package factory

import (
	"fmt"
)

//Cloud interface
type Cloud interface {
	Deploy() string
}

//AWS
type AWS struct {
	Name string
}

//NewAWS, AWS's Constructor
func NewAWS(name string) Cloud {
	return &AWS{Name: name}
}

//Deploy
func (a *AWS) Deploy() string {
	return fmt.Sprintf("Deploy using %s", a.Name)
}

//Azure
type Azure struct {
	Name string
}

//NewAzure, Azure's Constructor
func NewAzure(name string) Cloud {
	return &Azure{Name: name}
}

//Deploy
func (a *Azure) Deploy() string {
	return fmt.Sprintf("Deploy using %s", a.Name)
}

//GoogleCloud
type GoogleCloud struct {
	Name string
}

//NewGoogleCloud, GoogleCloud's Constructor
func NewGoogleCloud(name string) Cloud {
	return &GoogleCloud{Name: name}
}

//Deploy
func (g *GoogleCloud) Deploy() string {
	return fmt.Sprintf("Deploy using %s", g.Name)
}
