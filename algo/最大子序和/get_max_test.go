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
