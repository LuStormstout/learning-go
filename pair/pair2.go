/*
变量在持续赋值的过程中，它的 pair 是保持始终不变的。
*/
package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// Book 具体的类
type Book struct {
}

func (book *Book) ReadBook() {
	fmt.Println("read a book")
}

func (book *Book) WriteBook() {
	fmt.Println("write a book")
}

func main() {
	// b: pair<type:Book, value:Book{}地址>
	b := &Book{}

	// r: pair<type: , value: >
	var r Reader
	// r: pair<type:Book, value:Book{}地址>
	r = b
	r.ReadBook()

	// w: pair<type: , value: >
	var w Writer
	// w: pair<type:Book, value:Book{}地址>
	w = r.(Writer) // 此处的断言为什么会成功，是因为 w 和 r 具体的 type 是一致的
	w.WriteBook()

}
