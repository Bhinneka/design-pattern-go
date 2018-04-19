package factory

import (
	"fmt"
)

//Travis
type Travis struct {
	Name string
}

//NewTravis, Travis's Constructor
func NewTravis(name string) CITool {
	return &Travis{Name: name}
}

//Build
func (t *Travis) Build() string {
	return fmt.Sprintf("Build using %s", t.Name)
}

//Jenkins
type Jenkins struct {
	Name string
}

//NewJenkins, Jenkins's Constructor
func NewJenkins(name string) CITool {
	return &Jenkins{Name: name}
}

//Build
func (j *Jenkins) Build() string {
	return fmt.Sprintf("Build using %s", j.Name)
}

//Circle
type Circle struct {
	Name string
}

//NewCircle, Circle's Constructor
func NewCircle(name string) CITool {
	return &Circle{Name: name}
}

//Build
func (c *Circle) Build() string {
	return fmt.Sprintf("Build using %s", c.Name)
}
