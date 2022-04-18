package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (user *User) Call() {
	fmt.Println("user is called ... ")
	fmt.Printf("%v\n", user)
}

func DoFiledAndMethod(input interface{}) {
	// 获取 input 的 type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is :", inputType.Name())

	// 获取 input 的值
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is :", inputValue)

	// 通过 type 获取 User 里面的字段
	// 1、获取 interface 的 reflect.TypeOf （对应的数据类型），通过 Type 得到 NumField （User 总共有多少个字段），进行遍历
	// 2、得到每个 filed，（数据类型）
	// 3、通过 filed 有一个 interface() 方法得到对应的 value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 通过 type 获取 User 里面的方法
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

func main() {
	user := User{
		Id:   1,
		Name: "ZhangSan",
		Age:  18,
	}

	DoFiledAndMethod(user)
}
