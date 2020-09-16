package fib

// squares返回一个匿名函数。
// 该匿名函数每次被调用时都会返回下一个数的平方。
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}


func Fib(n int ) int {
	if n < 2 {
		return 1
	}
	return Fib(n-2) + Fib(n-1)
}


func Fib2(n int) int {
	var f = [3]int{1, 1, 2}
	for i := 2; i <= n; i++ {
		f[2] = f[0] + f[1]
		f[0] = f[1]
		f[1] = f[2]
	}
	return f[1]
}

func Fib3() func() int {
	var n,m int
	return func() int {
		if n <= 1 {
			n = 1
		}
		n,m = n+m,n
		return m
	}
}

