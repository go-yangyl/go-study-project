package cod

import "fmt"

func PrintNum(n int) {
	var max int

	for n != 0 {
		max = max*10 + 9
		n--
	}
	fmt.Println(max)
}
