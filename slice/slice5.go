/**
切片的截取
*/
package main

import "fmt"

func main() {
	s := []int{1, 2, 3} // len = 3, cap = 3

	// 打印子切片从索引 0 （包含）到索引 2（不包含）
	// 相当于打印切片的索引 0 和 1
	s1 := s[0:2] // [1 2]
	fmt.Println(s1)

	// 切片 s 和切片 s1 指向的是同一个内存地址
	// 任意修改其中一个切片，两个切片的值都会发生改变
	s1[0] = 100
	fmt.Println("s =", s)
	fmt.Println("s1 =", s1)

	fmt.Println("=========")
	s2 := make([]int, 3) // s2 = [0, 0, 0]
	// copy() 可以将底层数组的 slice 一起进行拷贝
	// 将 s 中的值依次拷贝到 s2 中
	// s 和 s2 各有自己的内存地址
	fmt.Println(s2)
	copy(s2, s)
	fmt.Println(s2)
}
