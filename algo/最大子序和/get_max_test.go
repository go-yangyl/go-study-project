package get_max

import (
	"fmt"
	"testing"
)

func TestGetSliceSortMaxNum(t *testing.T) {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	data := GetSliceSortMaxNum(arr)
	fmt.Println(data)
}

func TestGetStringMaxLength(t *testing.T) {
	s := "cdabad"
	begin, max := GetStringMaxLength(s)
	fmt.Println(begin, max)

}

func TestFindDisappearedNumbers(t *testing.T) {
	var s = "We are happy."

	fmt.Println(FindDisappearedNumbers(s))
}
