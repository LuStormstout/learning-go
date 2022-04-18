/*
管道（channel）的定义
	make(chan Type) 等阶于 make(chan Type, 0)
	make(chan Type, capacity)

channel <- value 发送 value 到 channel
<-channel 接受并将其丢弃
x := <-channel 从 channel 中接收数据，并赋值给 x
x, ok := <-channel 功能同上，同时检查通道是否已关闭或者是否为空
*/
package main

import "fmt"

func main() {
	// 定义一个 channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine 结束") // defer 是在 num := <-c 之后执行的
		fmt.Println("goroutine 正在运行 ... ")

		c <- 666 // 将 666 发送给 c
	}()

	num := <-c // 从 c 中接收数据，并赋值给 num

	fmt.Println("num =", num)
	fmt.Println("main goroutine 结束")
}
