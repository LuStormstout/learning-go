package route

import (
	"github.com/gorilla/mux"
	"goblog/pkg/routes"
	"net/http"
)

// Router 路由对象
var Router *mux.Router

// Initialize 初始化路由
func Initialize() {
	Router = mux.NewRouter()
	routes.RegisterWebRoutes(Router)
}

// Name2URL 通过路由名称来获取 URL
func Name2URL(routeName string, pairs ...string) string {
	url, err := Router.Get(routeName).URL(pairs...)
	if err != nil {
		return ""
	}

	return url.String()
}

// GetRouteVariable 获取 URI 路由参数
func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}
