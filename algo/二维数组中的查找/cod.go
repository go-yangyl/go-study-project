package cod

func FindArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	if len(matrix[0])  == 0 {
		return false
	}

	raws := len(matrix)
	cols := len(matrix[0]) -1

	raw := 0

	for raws > raw && cols >= 0 {
		if target == matrix[raw][cols] {
			return true
		}
		if target < matrix[raw][cols] {
			cols--
			continue
		}

		if target > matrix[raw][cols] {
			raw++
		}
	}
	return false
}
