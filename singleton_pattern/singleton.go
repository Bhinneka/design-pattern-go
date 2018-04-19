package singleton

import (
	"fmt"
	"sync"
)

//single
type bhinneka struct {
	message string
}

//ShowMessage
func (b *bhinneka) ShowMessage() string {
	return fmt.Sprintf("Message : %s", b.message)
}

var (
	instance *bhinneka
	once     sync.Once
)

//GetInstance
//return only bhinneka's single object instance
func GetInstance() *bhinneka {
	once.Do(func() {
		instance = &bhinneka{"from Bhinneka"}
	})
	return instance
}
