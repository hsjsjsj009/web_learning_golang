package framework

import "web_learning_golang/framework/url"

type app struct {
	listRoute [][]string
	allHandler []*handler
}

func NewApp() *app{
	return &app{listRoute: [][]string{},allHandler: []*handler{}}
}

//func (a *app) MainRoute(path string) *route{
//
//}

func (a *app) CreateHandler(path string) *handler{
	handle := &handler{path: path}
	a.allHandler = append(a.allHandler,handle)
	pathSlice := url.PathToSlice(path)
	a.listRoute = append(a.listRoute,pathSlice)
	return handle
}
