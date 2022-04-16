package main

import "fmt"

// const 来定义枚举类型
// 可以添在 const() 添加一个关键字 iota，每行的 iota 都会累加 1，第一行的 iota 的值默认是 0
const (
	BEIJING  = iota // iota = 0
	SHANGHAI        // iota = 1
	SHENZHEN        // iota = 2
)

// iota 只能够配合 const() 一起使用，iota 只有在 const 中进行累加效果
const (
	a, b = iota + 1, iota + 2 // iota = 0, a = iota + 1, b = iota + 2, a = 1, b = 2
	c, d                      // iota = 1, c = iota + 1, d = iota + 2, c = 2, d = 3
	e, f                      // iota = 2, e = iota + 1, f = iota + 2, e = 3, f = 4

	g, h = iota * 2, iota * 3 // iota = 3, g = iota * 2, h = iota * 3, g = 6, h = 9
	i, k                      // iota = 4, i = iota * 2, k = iota * 3, i = 8, k = 12
)

func main() {
	// 常量（只读属性）
	const length int = 10

	fmt.Println("length =", length)

	fmt.Println("BEIJING =", BEIJING)
	fmt.Println("BEIJING =", SHANGHAI)
	fmt.Println("BEIJING =", SHENZHEN)

	fmt.Println("a =", a, "b =", b)
	fmt.Println("c =", c, "d =", d)
	fmt.Println("e =", e, "f =", f)
	fmt.Println("g =", g, "h =", h)
	fmt.Println("i =", i, "k =", k)
}
