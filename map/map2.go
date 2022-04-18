/**
map 的基本操作
*/
package main

import "fmt"

// map 作为形参传递
func printMap(cityMap map[string]string) {
	// cityMap 引用传递（传的是原 cityMap 的指针过来的）
	for key, value := range cityMap {
		fmt.Println("key =", key)
		fmt.Println("value =", value)
	}
}

func ChangeValue(cityMap map[string]string) {
	cityMap["england"] = "London"
}

func main() {
	cityMap := make(map[string]string)

	// 添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	// 遍历
	printMap(cityMap)

	fmt.Println("============")

	// 删除
	delete(cityMap, "China")

	// 修改
	cityMap["USA"] = "DC"
	ChangeValue(cityMap)

	printMap(cityMap)
}
