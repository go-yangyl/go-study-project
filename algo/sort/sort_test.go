package sort

import (
	"fmt"
	"testing"
)

func TestMaoPao(t *testing.T) {
	var arr = []int{3, 2, 9, 1, 4, 5, 12, 10}

	//fmt.Println(MaoPao(arr))

	fmt.Println(XuanZhe(arr))

	KuaiPai()
}
