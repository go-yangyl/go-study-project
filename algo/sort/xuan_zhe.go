package sort

func XuanZhe(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		var min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}

	return arr

}
