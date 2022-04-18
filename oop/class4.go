/*
interface 通用万能类型
	interface{} 空接口，是一个万能数据类型
*/
package main

import (
	"fmt"
)

// myFunc interface{} 是一个万能类型，对应的 arg 可以传递任何类型的参数进来
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	// interface{} 怎么知道 arg 究竟是什么类型呢，此时应该怎么区分底层引用的数据类型呢
	// Go 给 interface{} 提供了"类型断言"的机制，用来判断当前万能数据类型底层究竟引用的是哪种数据类型
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg is string type, value =", value)
		fmt.Printf("value type is %T\n", value)
	}

	fmt.Println("====== 分割线 ======")
}

type Song struct {
	singer string
}

func main() {
	song := Song{singer: "Li"}

	myFunc(song)
	myFunc("Zhang")
	myFunc(100)
	myFunc(3.14)
	myFunc(true)
}
