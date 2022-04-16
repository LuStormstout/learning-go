/**
固定长度的数组
*/
package main

import "fmt"

func printArray(myArray [4]int) {
	// 值拷贝，将传入的 myArray3 的值拷贝一份到 myArray 中

	for index, value := range myArray {
		fmt.Println("index =", index, "value =", value)
	}

	// 值拷贝，原来的 myArray3 中的值并不会改变
	myArray[0] = 111
}

func main() {
	// 固定长度的数组
	var myArray1 [10]int
	myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := [4]int{11, 22, 33, 44}

	for i := 0; i < len(myArray1); i++ {
		fmt.Println(myArray1[i])
	}

	for index, value := range myArray2 {
		fmt.Println("index =", index, "value =", value)
	}

	// 查看数组的数据类型
	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)
	fmt.Printf("myArray3 types = %T\n", myArray3)
	/**
	在 Go 中 [10]int 和 [4]int 是不同的数据类型
	在下面的 printArray() 方法中如果传入的是 myArray1 则会报错
	输出结果
	myArray1 types = [10]int
	myArray2 types = [10]int
	myArray3 types = [4]int
	*/

	printArray(myArray3)
	fmt.Println("======= 在执行完 printArray() 之后 myArray3 中的值并没有发生改变 =======")
	for index, value := range myArray3 {
		fmt.Println("index =", index, "value =", value)
	}
}
