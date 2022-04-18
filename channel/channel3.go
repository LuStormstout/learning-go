/*
关闭 channel
*/
package main

import "fmt"

func main() {
	c := make(chan int) // 创建一个没有缓冲的 channel

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		// close 可以关闭一个 channel
		close(c)
	}()

	for {
		// ok 如果为 true 表示 channel 没有关闭，如果为 false 表示 channel 已经关闭
		// 先执行表达式 data, ok := <-c; 再判断 ok
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Printf("main finished ... \n")
}
