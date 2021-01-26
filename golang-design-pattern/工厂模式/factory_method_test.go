package factory_method

import (
	"fmt"
	"testing"
)

func TestFactoryOne_Get(t *testing.T) {
	var one = FactoryOne{}
	compute(one)

	var two = FactoryTwo{}
	compute(two)
}

func compute(factory InitFactory) {
	op := factory.Create()
	fmt.Println(op.Get())
}
