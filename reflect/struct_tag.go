/*
反射解析结构体标签
*/
package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name  string `info:"name" doc:"我的姓名"` // 结构体属性标签（起说明介绍的作用）
	Email string `info:"email" doc:"我有邮箱"`
	Sex   string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info:", tagInfo, "doc:", tagDoc)
	}
}

func main() {
	var resume resume

	findTag(&resume)
}
