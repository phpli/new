package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3, 4} //直接初始化4个元素的切片
	ss := add(s1, 10, 12, 13)
	fmt.Println(ss)
	fmt.Printf("ss:%v,len:%d,cap:%d\n", ss, len(ss), cap(ss))
	ss2 := deleteSlice(ss, 2)
	fmt.Println(ss2)
	fmt.Printf("ss2:%v,len:%d,cap:%d\n", ss2, len(ss2), cap(ss2))

	s2 := make([]int, 3, 4) //创建一个包含3个元素，5的容量的切片
	fmt.Printf("s2:%v,len:%d,cap:%d\n", s2, len(s2), cap(s2))

	//添加一个元素
	s2 = append(s2, 7)
	fmt.Printf("s2 first append:%v,len:%d,cap:%d\n", s2, len(s2), cap(s2))

	s2 = append(s2, 10)
	fmt.Printf("s2 sencond append:%v,len:%d,cap:%d\n", s2, len(s2), cap(s2))

	s3 := make([]int, 5)
	//等价：make([]int,5,5)
	fmt.Printf("s3:%v,len:%d,cap:%d\n", s3, len(s3), cap(s3))
	//ShareSlice()
}

func ShareSlice() {
	s1 := []int{1, 2, 3, 4}
	s2 := s1[2:]
	fmt.Printf("s1:%v,len:%d,cap:%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2:%v,len:%d,cap:%d\n", s2, len(s2), cap(s2))

	s2[0] = 99
	fmt.Printf("s1:%v,len:%d,cap:%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2:%v,len:%d,cap:%d\n", s2, len(s2), cap(s2))

	s2 = append(s2, 199)
	fmt.Printf("s1:%v,len:%d,cap:%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2:%v,len:%d,cap:%d\n", s2, len(s2), cap(s2))

	s2[1] = 1999
	fmt.Printf("s1:%v,len:%d,cap:%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2:%v,len:%d,cap:%d\n", s2, len(s2), cap(s2))
}

func add(slice []int, elems ...int) []int {
	return append(slice, elems...)
}

func deleteSlice(slice []int, index int) []int {
	if index < 0 || index > len(slice)-1 {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
