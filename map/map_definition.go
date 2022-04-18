/**
map 的定义
*/
package main

import "fmt"

func main() {
	// 1、第一种方式，声明 myMap1 是一种 map 类型，key 是 string，value 是 string
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 是一个空 map")
	}

	// 在使用 map 前，需要先用 make() 给 map 分配数据空间
	myMap1 = make(map[string]string, 10)

	myMap1["one"] = "c++"
	myMap1["two"] = "python"
	myMap1["three"] = "php"

	fmt.Println(myMap1)

	// 2、第二种方式
	myMap2 := make(map[int]string)
	myMap2[1] = "chengdu"
	myMap2[2] = "beijing"
	myMap2[3] = "guangzhou"
	fmt.Println(myMap2)

	// 3、第三种方式
	myMap3 := map[string]string{
		"one":   "美国",
		"two":   "日本",
		"three": "英国",
	}
	fmt.Println(myMap3)
}
