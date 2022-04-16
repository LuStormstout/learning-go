package main

import "fmt"

func swap1(a int, b int) {
	var temp int
	temp = a
	a = b
	b = temp

	fmt.Println("swap1() 方法中：a =", a, "b =", b)
}

func swap2(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

func main() {
	var a int = 10
	var b int = 20

	swap1(a, b)
	fmt.Println("main() 方法中执行 swap1() 之后：a =", a, "b =", b)

	swap2(&a, &b)
	fmt.Println("main() 方法中执行 swap2() 之后：a =", a, "b =", b)

	var p *int
	p = &a
	fmt.Println(&a)
	fmt.Println(p)

	var pp **int // 二级指针
	pp = &p
	fmt.Println(&p)
	fmt.Println(pp)
}
