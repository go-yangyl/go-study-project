package sort

func KuaiPai() {
	QuickSort(0, len(li)-1)
}

var li = []int{3, 2, 9, 1, 4, 5, 12, 10}

func QuickSort(left, right int) {
	if left >= right {
		return
	}

	i := left
	j := right

	temp := li[left]

	for i != j {
		for li[j] >= temp && j > i {
			j--
		}

		for li[i] <= temp && i < j {
			i++
		}

		if i < j {
			li[i], li[j] = li[j], li[i]
		}

	}

	li[left] = li[i]
	li[i] = temp

	QuickSort(left, i-1)  //继续处理左边的，这里是一个递归的过程
	QuickSort(i+1, right) //继续处理右边的 ，这里是一个递归的过程
}
