package find

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 5, 7, 15, 25, 30, 36, 39, 51, 67, 78, 80, 82, 85, 91, 92, 97}
	BinarySearch(arr, 0, len(arr), 30)
}
