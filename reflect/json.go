/*
结构体标签在 json 中的应用
*/
package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title,omitempty"`
	Year   int      `json:"year,omitempty"`
	Price  int      `json:"price,omitempty"`
	Actors []string `json:"actors,omitempty"`
}

func main() {
	movie := Movie{
		Title:  "无间道",
		Year:   2002,
		Price:  25,
		Actors: []string{"刘德华", "梁朝伟", "黄秋生", "曾志伟"},
	}

	// 编码的过程 结构体 -> json （结构体转换为 json）
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)

	// 解码的过程 json -> 结构体 （json 转换为结构体）
	// jsonStr = {"title":"无间道","year":2002,"price":25,"actors":["刘德华","梁朝伟","黄秋生","曾志伟"]}
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error", err)
		return
	}
	fmt.Printf("%v\n", myMovie)
}
