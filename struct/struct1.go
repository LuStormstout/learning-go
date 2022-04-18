package main

import "fmt"

// 声明一种新的数据类型 myInt，相当于 int 的一个别名
type myInt int

// Book 定义一个结构体
type Book struct {
	title string
	auth  string
}

func changeBook1(book Book) {
	// 传递的是一个 book 副本
	// 值并不会被改变
	book.auth = "LiSi"
}

func changeBook2(book *Book) {
	// 指针传递
	// 值会发生改变
	book.auth = "WangMaZi"
}

func main() {
	var a myInt = 10
	fmt.Println("a =", a)
	fmt.Printf("type of a = %T\n", a)

	var book1 Book
	book1.title = "Golang"
	book1.auth = "ZhangSanFen"
	fmt.Printf("%v\n", book1)

	changeBook1(book1)
	fmt.Printf("%v\n", book1)

	changeBook2(&book1)
	fmt.Printf("%v\n", book1)
}
