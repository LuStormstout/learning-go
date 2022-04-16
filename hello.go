package main // 程序包的名

import (
	"fmt"
	"time"
)

// main 函数
func main() { // 函数的 { 一定是和函数名在同一行的，否则编译报错
	// golang 中的表达式，加";"和不加都可以，建议不加
	fmt.Println("hello Go!")

	time.Sleep(1 * time.Second)
}
