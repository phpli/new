package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3}
	fmt.Printf("a1:%v, \n cap:%d,\n len:%d\n", a1, cap(a1), len(a1))
	fmt.Printf("a1[1]:%d\n", a1[1])

	//a2 := [3]int{}
	var a2 [3]int
	fmt.Printf("a2:%v, \n cap:%d,\n len:%d\n", a2, cap(a2), len(a2))
	//fib := fibonaccGenerator()
	for i := 0; i < 10; i++ {
		//fmt.Printf("fibonacc[%d]: %v\n", i, fib())
		//fmt.Printf("fibonacc[%d]\n", iffibonacc(i))
		fmt.Printf("fibonacc[%d]\n", fibonacci(i))
	}
}

// 闭包形式
func fibonaccGenerator() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// 迭代方式
func iffibonacc(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

//递归方式

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
