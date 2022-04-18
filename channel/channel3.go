/*
关闭 channel

向已经关闭的 channel 发送数据会报错
channel 已经关闭，但是仍然是可以从 channel 冲读取数据的
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

	/*
		for {
			// ok 如果为 true 表示 channel 没有关闭，如果为 false 表示 channel 已经关闭
			// 先执行表达式 data, ok := <-c; 再判断 ok
			if data, ok := <-c; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
	*/

	// 以上可以用 range 关键字简写
	// 用 range 来迭代不断操作 channel
	for data := range c {
		fmt.Println(data)
	}

	fmt.Printf("main finished ... \n")
}
