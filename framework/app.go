package framework

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"web_learning_golang/framework/response"
	"web_learning_golang/framework/url"
)

type app struct {
	listRoute [][]string
	allHandler []*handler
}

func NewApp() *app{
	return &app{listRoute: [][]string{},allHandler: []*handler{}}
}

func (a *app) MainRoute(path string) *route{
	return &route{a,path}
}

func (a *app) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	var (
		responseObj response.Response
		statusCode int8
	)
	path := r.URL.Path
	method := r.Method
	pathVar,found ,idx:= url.ParsePath(path,a.listRoute)
	if !found {
		w.WriteHeader(404)
		_, _ = w.Write([]byte("404 Not Found"))
		log.Println(fmt.Sprintf("[%s] %s status code %d",strings.ToUpper(method),path,404))
		return
	}

	requestObj := newRequest(pathVar,r)
	handler := a.allHandler[idx]

	if handler.allMethod != nil {
		responseObj = handler.allMethod(requestObj)
	}else if handler.getMethod != nil && method == "GET" {
		responseObj = handler.getMethod(requestObj)
	}else if handler.postMethod != nil && method == "POST" {
		responseObj = handler.postMethod(requestObj)
	}else if handler.deleteMethod != nil && method == "DELETE" {
		responseObj = handler.deleteMethod(requestObj)
	}else if handler.putMethod != nil && method == "PUT" {
		responseObj = handler.putMethod(requestObj)
	}

	if responseObj != nil {
		responseObj.WriteHeader(w)
		err := responseObj.SendResponse(w)
		statusCode = http.StatusOK
		if err != nil {
			statusCode = http.StatusInternalServerError
			http.Error(w,err.Error(), int(statusCode))
			log.Fatal(fmt.Sprintf("Error %s\n",err.Error()))
			return
		}
		log.Println(fmt.Sprintf("[%s] %s status code %d",method,path,statusCode))
		return
	}

	log.Println(fmt.Sprintf("[%s] %s status code %d",method,path,http.StatusNotFound))
	w.WriteHeader(404)
	_, _ = w.Write([]byte(fmt.Sprintf("Path %s doesn't accept method %s",path,r.Method)))

}

func (a *app) Start(url string) {
	log.Println(fmt.Sprintf("Server run in localhost port %s",url))
	log.Fatal(http.ListenAndServe(url, a))
}
