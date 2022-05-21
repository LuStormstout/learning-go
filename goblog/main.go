package main

import (
	"goblog/bootstrap"
	"goblog/http/middlewares"
	"goblog/pkg/logger"
	"net/http"
)

func main() {
	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
