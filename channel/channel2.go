/*
在 channel 未关闭的情况下
	当 channel 已经满了的时候，再向里面写数据，就会阻塞
	当 channel 为空，从里面取数据也会阻塞
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) // 带有缓冲的 channel

	fmt.Println("len(c) =", len(c), "cap(c) =", cap(c))

	go func() {
		defer fmt.Println("子 go 程结束")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子 go 程正在执行，发送的元素 =", i, "len(c) =", len(c), "cap(c) =", cap(c))
		}
	}()

	// 为了更好地看到效果，让程序休眠 2 秒
	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c // 从 c channel 中接收数据，并赋值给 num
		fmt.Println("num =", num)
	}

	fmt.Println("main 结束")
}
