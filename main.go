package main

import (
	"web_learning_golang/framework"
	"web_learning_golang/framework/response"
)

func main(){
	app := framework.NewApp()

	mainRoute := app.MainRoute("/")
	mainHandler := mainRoute.CreateHandler()
	mainHandler.Get(func(request *framework.HttpRequest) response.Response {
		return response.NewJsonResponse(map[string]interface{}{})
	})

	app.Start(":8080")
}
