/**
defer 的执行顺序
*/
package main

import "fmt"

func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}

func main() {
	defer func1()
	defer func2()
	defer func3()

	/**
	执行结果（压栈原理，先进后出的顺序）
	➜  defer go run defer2.go
	C
	B
	A
	*/
}
