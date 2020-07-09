package main

import (
	"web_learning_golang/framework"
	"web_learning_golang/framework/response"
)

func main(){
	app := framework.NewApp()

	mainRoute := app.MainRoute("/")
	mainRoute.Get(func(request *framework.HttpRequest) response.Response {
		return response.Redirect("/asd/asdasdsa/",request.Request)
	})

	mainRoute.Post(func(request *framework.HttpRequest) response.Response {
		_ = request.Request.ParseForm()
		return response.JsonResponse(map[string]interface{}{
			"form" : request.Request.Form["asdasdsa"],
		})
	})

	secondRoute := app.MainRoute("/asd/:asd")
	secondRoute.Get(func(request *framework.HttpRequest) response.Response {
			return response.JsonResponse(map[string]interface{}{
				"asd":request.GetVariablePath("asd"),
			})
	})

	app.Start(":8081")
}
