/**
defer 和 return 谁先谁后
*/
package main

import "fmt"

func deferFunc() int {
	fmt.Println("defer func called ...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called ...")
	return 0
}

func deferOrReturn() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	deferOrReturn()

	/**
	执行结果
	defer 是在一个方法的生命周期结束之后再被调用
	结果是 return 会执行在 defer 之前
	➜  defer go run defer3.go
	return func called ...
	defer func called ...
	*/
}
