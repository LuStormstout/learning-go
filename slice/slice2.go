/**
动态数组，切片 slice
*/
package main

import "fmt"

func printSlice(myArray []int) {
	// 引用传递（传递的是当前数组「myArray」的指针进来）

	// _ 表示匿名的变量（访问不了，处理一些并不想去关心的内容是就可以这么写了）
	for _, value := range myArray {
		fmt.Println("value =", value)
	}

	// 修改之后原 myArray 中的值也会变更
	myArray[0] = 100
}

func main() {
	myArray := []int{1, 2, 3, 4} // 动态数组，切片 slice

	fmt.Printf("myArray type is %T\n", myArray)
	/**
	输出结果
	➜  slice git:(main) ✗ go run slice2.go
	myArray type is []int
	*/

	printSlice(myArray)

	fmt.Println("======= 执行完 printSlice() 之后 myArray 中的值发生了改变 =======")
	for _, value := range myArray {
		fmt.Println("value =", value)
	}
}
