package find

import "fmt"

func BinarySearch(arr []int, leftIndex, rightIndex, findVal int) {
	if leftIndex > rightIndex {
		fmt.Println("没有找到")
	}

	for leftIndex != rightIndex {
		middle := (leftIndex + rightIndex) / 2

		if arr[middle] == findVal {
			fmt.Println(middle)
			break
		} else if findVal > arr[middle] {
			leftIndex = middle + 1
		} else if findVal < arr[middle] {
			rightIndex = middle - 1
		}
	}
}
