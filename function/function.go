package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("========= foo1 =========")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	c := 100
	return c
}

// 返回多个返回值，匿名的
func foo2(a string, b int) (int, int) {
	fmt.Println("========= foo2 =========")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	return 2021, 2022
}

// 返回多个返回值，有形参名称的返回值 r1 r2
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("========= foo3 =========")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// r1 r2 属于 foo3 的形参，初始化默认的值是 0
	// r1 r2 作用域是 foo3 内
	fmt.Println("r1 =", r1)
	fmt.Println("r2 =", r2)

	// 给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000

	return
}

// 返回值的类型是一样的话可以写在一起 int
func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("========= foo4 =========")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// 给有名称的返回值变量赋值
	r1 = 714
	r2 = 147

	return
}

func main() {
	c := foo1("cola", 500)
	fmt.Println("c =", c)

	ret1, ret2 := foo2("beer", 350)
	fmt.Println("ret1 =", ret1, "ret2 =", ret2)

	ret1, ret2 = foo3("phone", 10999)
	fmt.Println("ret1 =", ret1, "ret2 =", ret2)

	ret1, ret2 = foo4("car", 2018)
	fmt.Println("ret1 =", ret1, "ret2 =", ret2)
}
