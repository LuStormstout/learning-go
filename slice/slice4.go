/**
slice 使用方式
追加
*/
package main

import "fmt"

// 切片的扩容机制，再往一个切片中 append 的时候，如果长度增加后超过当前切片容量，则将容量增加到当前容量的 2 倍
func main() {
	// 该切片的长度是 3，容量是 5。
	var numbers1 = make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers1), cap(numbers1), numbers1)

	// 向 numbers1 切片中追加一个元素 1，该切片的 len = 4，[0, 0, 0, 1]，cap = 5
	numbers1 = append(numbers1, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers1), cap(numbers1), numbers1)

	// 向 numbers1 切片中追加一个元素 2，该切片的 len = 5，[0, 0, 0, 1, 2]，cap = 5
	numbers1 = append(numbers1, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers1), cap(numbers1), numbers1)

	// 向 numbers1 切片中追加一个元素 3，此时原来的五个容量已经用完了，Go 会自动将容量帮我们扩大到 10，原来声明切片 numbers1 的时候的容量 5 就会成为扩大容量的"步长"
	// 此时 numbers1 切片中的 len = 6，[0, 0, 0, 1, 2, 3]，cap = 10
	numbers1 = append(numbers1, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers1), cap(numbers1), numbers1)

	/**
	输出结果
	len = 3, cap = 5, slice = [0 0 0]
	len = 4, cap = 5, slice = [0 0 0 1]
	len = 5, cap = 5, slice = [0 0 0 1 2]
	len = 6, cap = 10, slice = [0 0 0 1 2 3]
	*/

	fmt.Println("============ 声明切片的时候不去指定容量 cap ============")
	// 声明切片 numbers2 的时候没有指定容量，默认的容量就会是长度的大小，
	// len = 3, cap = 3, slice = [0 0 0]
	var numbers2 = make([]int, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)
	// 在往声明切片没有指定容量的切片 numbers2 中追加元素的时候，numbers2 的容量会扩大到 6，扩大的"步长"为当前切片的容量 3，再下次扩大容量的"步长"就为 6
	// len = 4, cap = 6, slice = [0 0 0 1]
	numbers2 = append(numbers2, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)
	numbers2 = append(numbers2, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)
	numbers2 = append(numbers2, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)
	numbers2 = append(numbers2, 4)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)
}
