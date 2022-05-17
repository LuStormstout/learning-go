package routes

import (
	"github.com/gorilla/mux"
	"goblog/app/http/controllers"
	"net/http"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(r *mux.Router) {
	// 静态页面
	pagesController := new(controllers.PagesController)
	r.HandleFunc("/", pagesController.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pagesController.About).Methods("GET").Name("about")
	r.NotFoundHandler = http.HandlerFunc(pagesController.NotFound)

	// 文章相关页面
	articlesController := new(controllers.ArticlesController)
	r.HandleFunc("/articles/{id:[0-9]+}", articlesController.Show).Methods("GET").Name("article.show")
	r.HandleFunc("/articles", articlesController.Index).Methods("GET").Name("articles.index")
	r.HandleFunc("/articles", articlesController.Store).Methods("POST").Name("articles.store")
	r.HandleFunc("/articles/create", articlesController.Create).Methods("GET").Name("articles.create")
}
