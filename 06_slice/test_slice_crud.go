package main

import "fmt"

func main() {
	// 增加
	// var s1 = []int{}
	// s1 = append(s1, 1)
	// s1 = append(s1, 2)
	// fmt.Printf("s1: %v\n", s1)

	// 删除
	var s1 = []int{1, 2, 3, 4}
	var s2 = append(s1[:2], s1[3:]...)
	// append(s[:index], s[index+1:]...)
	fmt.Printf("s2: %v\n", s2)

	// 拷贝
	// var s1 = []int{1, 2, 3, 4}
	// var s2 = make([]int, len(s1))
	// copy(s2, s1)
	// fmt.Printf("s2: %v\n", s2)
}
