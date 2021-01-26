package facade

import (
	"fmt"
	"testing"
)

func TestNewFacade(t *testing.T) {
	one := &FacadeOne{Name: "yangyl"}

	facade := NewFacade()
	facade.One = one
	fmt.Println(facade.One.Get())
}
