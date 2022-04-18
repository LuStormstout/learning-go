/*
反射
*/
/*
func ValueOf(i interface{}) Value {...}
ValueOf 用来获取输入参数接口中数据的值，如果接口为空则返回 0

func TypeOf(i interface{}) Type {...}
TypeOf 用来动态获取出入参数接口中的值的类型，如果接口为空则返回 nil
*/

package main

import (
	"fmt"
	"reflect"
)

func reflectNumber(arg interface{}) {
	fmt.Println("type :", reflect.TypeOf(arg))
	fmt.Println("value :", reflect.ValueOf(arg))
}

func main() {
	var num float64 = 1.2345

	reflectNumber(num)
}
