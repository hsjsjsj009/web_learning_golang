package main

import (
	"web_learning_golang/framework"
)

func main(){
	app := framework.NewApp()

	mainRoute := app.MainRoute("/")
	//mainHandler := mainRoute.CreateHandler()
	//mainHandler.Get(func(request *framework.HttpRequest) response.Response {
	//	return response.JsonResponse(map[string]interface{}{}).
	//		SetStatusCode(200).
	//		AddHeader("test","asd")
	//})

	app.Start(":8080")
}
