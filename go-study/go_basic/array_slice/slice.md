## slice

```go
    1. 切片：切片是数组的一个引用，底层仍然是数组，因此切片是引用类型
    2. append改变长度，会产生新的内存空间，不会改变原来的底层数组,只有修改切片的值，原来的切片和数字也会发生改变
    3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。 
    4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
    5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int。
    6. 如果 slice == nil，那么 len、cap 结果都等于 0。
```


```go
func make([]T, len, cap) []T
```

第一个参数是 []T，T 即元素类型，第二个参数是长度 len，即初始化的切片拥有多少个元素，第三个参数是容量 cap，容量是可选参数，默认等于长度。使用内置函数 len 和 cap 可以得到切片的长度和容量，例如：

```go
func printLenCap(nums []int) {
	fmt.Printf("len: %d, cap: %d %v\n", len(nums), cap(nums), nums)
}

func TestSliceLenAndCap(t *testing.T) {
	nums := []int{1}
	printLenCap(nums) // len: 1, cap: 1 [1]
	nums = append(nums, 2)
	printLenCap(nums) // len: 2, cap: 2 [1 2]
	nums = append(nums, 3)
	printLenCap(nums) // len: 3, cap: 4 [1 2 3]
	nums = append(nums, 3)
	printLenCap(nums) // len: 4, cap: 4 [1 2 3 3]
}
```

容量是当前切片已经预分配的内存能够容纳的元素个数，如果往切片中不断地增加新的元素。如果超过了当前切片的容量，就需要分配新的内存，并将当前切片所有的元素拷贝到新的内存块上。因此为了减少内存的拷贝次数，容量在比较小的时候，一般是以 2 的倍数扩大的，例如 2 4 8 16 …，当达到 2048 时，会采取新的策略，避免申请内存过大，导致浪费。Go 语言源代码 runtime/slice.go 中是这么实现的，不同版本可能有所差异：

```go
newcap := old.cap
doublecap := newcap + newcap
if cap > doublecap {
	newcap = cap
} else {
	if old.len < 1024 {
		newcap = doublecap
	} else {
		// Check 0 < newcap to detect overflow
		// and prevent an infinite loop.
		for 0 < newcap && newcap < cap {
			newcap += newcap / 4
		}
		// Set newcap to the requested cap when
		// the newcap calculation overflowed.
		if newcap <= 0 {
			newcap = cap
		}
	}
}
```

切片和数组很相似，按照下标进行索引。切片本质是一个数组片段的描述，包括了数组的指针，这个片段的长度和容量(不改变内存分配情况下的最大长度

```go
struct {
    ptr *[]T
    len int
    cap int
}
```

```go
nums := make([]int, 0, 8)
nums = append(nums, 1, 2, 3, 4, 5)
nums2 := nums[2:4]
printLenCap(nums)  // len: 5, cap: 8 [1 2 3 4 5]
printLenCap(nums2) // len: 2, cap: 6 [3 4] 数组前面的地址空间无法使用

nums2 = append(nums2, 50, 60)
printLenCap(nums)  // len: 5, cap: 8 [1 2 3 4 50]
printLenCap(nums2) // len: 4, cap: 6 [3 4 50 60]
```

- nums2 执行了一个切片操作 `[2, 4)`，此时 nums 和 nums2 指向的是同一个数组。
- nums2 增加 2 个元素 50 和 60 后，将底层数组下标 [4] 的值改为了 50，下标[5] 的值置为 60。
- 因为 nums 和 nums2 指向的是同一个数组，因此 nums 被修改为 [1, 2, 3, 4, 50]。



#### append

- 当 append 之后的长度小于等于 cap，将会直接利用原底层数组剩余的空间。
- 当 append 后的长度大于 cap 时，则会分配一块更大的区域来容纳新的底层数组。



#### delete

切片的底层是数组，因此删除意味着后面的元素需要逐个向前移位。每次删除的复杂度为 O(N)，因此切片不合适大量随机删除的场景，这种场景下适合使用链表。



#### insert

```go
arr = append(arr[:i],append([]int{x},append(arr[i:]...)))
```

insert 和 append 类似。即在某个位置添加一个元素后，将该位置后面的元素再 append 回去。复杂度为 O(N)。因此，不适合大量随机插入的场景。



#### 切片的性能陷阱

在已有切片的基础上进行切片，不会创建新的底层数组。因为原来的底层数组没有发生变化，内存会一直占用，直到没有变量引用该数组。因此很可能出现这么一种情况，原切片由大量的元素构成，但是我们在原切片的基础上切片，虽然只使用了很小一段，但底层数组在内存中仍然占据了大量空间，得不到释放。比较推荐的做法，使用 `copy` 替代 `re-slice`。

```go
func lastNumsBySlice(origin []int) []int {
	return origin[len(origin)-2:]
}

func lastNumsByCopy(origin []int) []int {
	result := make([]int, 2)
	copy(result, origin[len(origin)-2:])  // 释放origin底层数组
	return result
}
```

- 第一个函数直接在原切片基础上进行切片。
- 第二个函数创建了一个新的切片，将 origin 的最后两个元素拷贝到新切片上，然后返回新切片。