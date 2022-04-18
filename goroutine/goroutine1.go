package main

import (
	"fmt"
	"time"
)

// 从/子 goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 主 goroutine
func main() {
	// 创建一个 go程 （goroutine）去执行 newTask() 流程
	go newTask()

	fmt.Println("main goroutine has exited")

	// 主的一直死循环，这样从/子的才能够执行，主的 goroutine 退出了，从/子的 goroutine 也就没有办法继续执行了
	/*
		i := 0
		for {
			i++
			fmt.Printf("main goroutine: i = %d\n ", i)
			time.Sleep(1 * time.Second)
		}
	*/

}
