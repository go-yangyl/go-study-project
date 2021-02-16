package get_max

import (
	"math"
)

func GetSliceSortMaxNum(nums []int) float64 {
	dp := make([]int, len(nums), len(nums))

	dp[0] = nums[0]
	ans := float64(nums[0])

	for i := 1; i < len(nums); i++ {
		dp[i] = int(math.Max(float64(dp[i-1]), float64(0)) + float64(nums[i]))
		ans = math.Max(ans, float64(dp[i]))
	}
	return ans
}

func GetStringMaxLength(str string) (int, int) {
	max := 1
	begin := 0

	dp := make([][]bool, len(str))
	for i := 0; i < len(str); i++ {
		dp[i] = make([]bool, len(str))
		dp[i][i] = true
	}

	for j := 1; j < len(str); j++ { // 遍历行
		for i := 0; i < j; i++ { // 遍历列
			if str[j] != str[i] {
				dp[i][j] = false
				continue
			}

			if j-i < 3 { // 字符串就 1 或 2个
				dp[i][j] = true
			} else {
				// 保障内部是回文字符串
				dp[i][j] = dp[i+1][j-1]
			}

			if dp[i][j] && j-i+1 > max {
				max = j - i + 1
				begin = i
			}
		}
	}
	return begin, max
}

func FindDisappearedNumbers(s string) string {
	var str = ""
	for _, v := range s {
		if string(v) == " " {
			str += "%20"
			continue
		}
		str += string(v)
	}

	return str

}
