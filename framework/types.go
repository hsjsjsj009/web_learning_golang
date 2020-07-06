package framework

import (
	"net/http"
)

type Route struct {
	Method int8
	Path string
	Handler func(r *http.Request) interface{}
}


