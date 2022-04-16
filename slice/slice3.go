/**
slice 声明方式
*/
package main

import "fmt"

func main() {
	// 1、声明 slice1 是一个切片，并且初始化，默认值是 1，2，3。长度 len 是 3
	//slice1 := []int{1, 2, 3}

	// 2、声明 slice 是一个切片，但是并没有给 slice 分配空间
	var slice1 []int
	//slice1 = make([]int, 3) // 开辟三个空间，默认值是 0

	// 3、声明 slice1 是一个切片，同时给 slice1 分配 3 个空间，初始化值是 0
	//var slice1 []int = make([]int, 3)

	// 4、声明 slice1 是一个切片，同时给 slice1 分配 3 个空间，初始化值是 0，通过 := 推导出 slice1 是一个切片
	//slice1 := make([]int, 3)

	//slice1[0] = 100
	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	// 判断一个 slice 是否为空（没有空间）
	if slice1 == nil {
		fmt.Println("slice1 是一个空切片")
	} else {
		fmt.Println("slice1 是有空间的")
	}
}
