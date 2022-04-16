package main

import "fmt"

func main() {
	// 写入 defer 关键字
	// 会在当前函数执行结束之前触发（类似于析构函数）
	defer fmt.Println("main end 1")
	defer fmt.Println("main end 2")

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")

	/**
	执行结果
	因为 "压栈" 所以先执行「main end 2」再执行「main end 1」
	入栈的时候依次顺序是「main end 1」入栈 ➡️「main end 2」入栈 ➡️ ...其他的代码（相当于 main end 2 压在 main end 1 上面）
	所以出栈的时候就会先执行压在上面的「main end 2」再执行「main end 1」（这样就便于理解了）
	先进后出的顺序
	➜  defer go run defer1.go
	main::hello go 1
	main::hello go 2
	main end 2
	main end 1
	*/

}
