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
