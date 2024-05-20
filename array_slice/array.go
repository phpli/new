package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3}
	fmt.Printf("a1:%v, \n cap:%d,\n len:%d\n", a1, cap(a1), len(a1))
	fmt.Printf("a1[1]:%d\n", a1[1])

	//a2 := [3]int{}
	var a2 [3]int
	fmt.Printf("a2:%v, \n cap:%d,\n len:%d\n", a2, cap(a2), len(a2))
}
