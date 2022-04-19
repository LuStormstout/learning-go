/*
channel 和 select
单流程下一个 go 只能监控一个 channel 的状态，select 可以完成监控多个 channel 的状态

select 可以在同一流程下监听多个 channel 的状态
select 具备监控多路 channel 状态的功能
	chan1 := make(chan int)
	chan2 := make(chan int)

	select {
	case <-chan1:
		// 如果 chan1 成功读取到数据，则执行该 case 处理语句
	case chan2 <- 1:
		// 如果成功向 chan2 写入数据，则执行该 case 处理语句
	default:
		// 如果上面的都没有执行成功，则执行 default 处理语句
	}
*/

package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1

	for {
		select {
		case c <- x:
			// 如果 c 可写，则该 case 就会进来
			x, y = y, x+y
		case <-quit: // quite channel 可读的话，说明那边写入 c channel 的循环完事儿了
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	// sub go
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	// main go
	fibonacci(c, quit)
}
