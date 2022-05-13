package controllers

import (
	"fmt"
	"net/http"
)

// PagesController 处理静态页面
type PagesController struct {
}

// Home 首页
func (*PagesController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>👋 Hello, this is a blogging practice project built in Go.</h1>")
}

// About 关于我们页面
func (*PagesController) About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "此博客是用以记录编程笔记，如你有反馈或建议，请联系"+
		"<a href=\"mailto:lustormstout@gmail.com\">lustormstout@gmail.com</a>")
}

// NotFound 404 页面
func (*PagesController) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 🙁 </h1>"+
		"如有疑惑，请联系我们。")
}
