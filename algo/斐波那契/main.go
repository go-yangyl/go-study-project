package main

/*
写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：

F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	n1 := 0
	n2 := 1
	sum := 0
	for i := 0; i <= n-2; i++ {
		sum = n1 + n2
		n1 = n2
		n2 = sum
	}
	return sum % 1000000007
}

func foo(arg_val int) *int {

	var foo_val1 int = 11
	var foo_val2 int = 12
	var foo_val3 int = 13
	var foo_val4 int = 14
	var foo_val5 int = 15

	//此处循环是防止go编译器将foo优化成inline(内联函数)
	//如果是内联函数，main调用foo将是原地展开，所以foo_val1-5相当于main作用域的变量
	//即使foo_val3发生逃逸，地址与其他也是连续的
	for i := 0; i < 5; i++ {
		println(&arg_val, &foo_val1, &foo_val2, &foo_val3, &foo_val4, &foo_val5)
	}

	//返回foo_val3给main函数
	return &foo_val3
}

func main() {
	var a = make([]*int, 1)
	b := 1
	a[0] = &b
}
