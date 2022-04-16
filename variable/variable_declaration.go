package main

import "fmt"

// 声明全剧变量 方法一、方法二、方法三是可以的
// := 只能够用在函数体内来声明
var gA int = 100
var gB = 200

// 四种变量的声明方式
func main() {
	// 方法一：声明一个变量，默认的值是 0
	var a int
	fmt.Println("a =", a)
	fmt.Printf("type of a = %T\n", a)

	// 方法二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b =", b)
	fmt.Printf("type of b = %T\n", b)
	var bb string = "computer"
	fmt.Printf("bb = %s, type of bb = %T\n", bb, bb)

	// 方法三：在初始化的时候，可以省区数据类型，通过自动匹配当前变量的数据类型
	var c = 100
	fmt.Println("c =", c)
	fmt.Printf("type of c = %T\n", c)
	var cc = "phone"
	fmt.Printf("cc = %s, type of cc = %T\n", cc, cc)

	// 方法四：（常用的方法）省区 variable 关键字，直接自动匹配（只能够用在函数体内来声明）
	e := 100
	fmt.Println("e =", e)
	fmt.Printf("type of e = %T\n", e)
	f := "cola"
	fmt.Printf("f = %s, type of f = %T\n", f, f)
	g := 3.14
	fmt.Println("g =", g)
	fmt.Printf("type of g = %T\n", g)

	// 打印全局变量
	fmt.Println("gA =", gA, "gB =", gB)

	// 声明多个变量
	var xx, yy int = 100, 200
	var kk, ll = 100, "cup"
	fmt.Println("xx =", xx, "yy =", yy)
	fmt.Println("kk =", kk, "ll =", ll)
	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv =", vv, "jj =", jj)
}
