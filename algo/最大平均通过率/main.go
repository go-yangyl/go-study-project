package main

import "container/heap"

func main() {

}

/*
解题思路：
一、先分别计算所有班级如果加一个人的通过变化率，放入优先队列
二、取队头的班级（加一个人后通过率提升最大的班级），该班级的通过人数和总人数各自+1
三、更新第二步的班级的新的通过变化率，随后将其重新加入优先队列
有多少人就重复第二和第三步多少次即可。
PS：Go实现优先队列要是也能只调用一个API就好了。

作者：thebug-v
链接：https://leetcode-cn.com/problems/maximum-average-pass-ratio/solution/goyou-xian-dui-lie-by-thebug-v-uxn2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

type maxHeap []Pair

type Pair struct {
	Index int
	C     float64
}

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i].C > h[j].C
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(num interface{}) {
	*h = append(*h, num.(Pair))
}
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	q := &maxHeap{}
	for i, v := range classes {
		a := float64(v[0]) / float64(v[1])
		b := float64(v[0]+1) / float64(v[1]+1)
		heap.Push(q, Pair{
			Index: i,
			C:     b - a,
		})
	}
	heap.Init(q)
	for i := 0; i < extraStudents; i++ {
		temp := heap.Pop(q).(Pair)
		classes[temp.Index][0]++
		classes[temp.Index][1]++
		a := float64(classes[temp.Index][0]) / float64(classes[temp.Index][1])
		b := float64(classes[temp.Index][0]+1) / float64(classes[temp.Index][1]+1)
		temp.C = b - a
		heap.Push(q, temp)
	}
	var sum float64
	for _, v := range classes {
		sum += float64(v[0]) / float64(v[1])
	}
	return sum / float64(len(classes))
}
