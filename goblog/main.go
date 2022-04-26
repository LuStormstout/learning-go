package main

import (
	"fmt"
	"net/http"
	"strings"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>👋 Hello, this is a blogging practice project built in Go.</h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>请求页面未找到 🙁 </h1>"+
			"如有疑惑，请联系我们。")
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "此博客是用以记录编程笔记，如你有反馈或建议，请联系"+
		"<a href=\"mailto:lustormstout@gmail.com\">lustormstout@gmail.com</a>")
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", defaultHandler)
	router.HandleFunc("/about", aboutHandler)

	// 文章详情
	router.HandleFunc("/articles/", func(writer http.ResponseWriter, request *http.Request) {
		id := strings.SplitN(request.URL.Path, "/", 3)[2]
		fmt.Fprint(writer, "文章 ID："+id)
	})

	// 列表 OR 创建
	router.HandleFunc("/articles", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			fmt.Fprint(writer, "访问文章列表")
		case "POST":
			fmt.Fprint(writer, "创建新的文章")
		}
	})
	http.ListenAndServe(":3000", router)
}
