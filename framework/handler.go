package framework

import (
	"net/http"
	"web_learning_golang/framework/response"
)

const GET = 1
const POST = 2
const PUT = 3
const DELETE = 4

type HttpRequest struct {
	request *http.Request
	pathVariable map[string]string
}

type handler struct {
	path string
	allMethod func(*HttpRequest) response.Response
	getMethod func(*HttpRequest) response.Response
	postMethod func(*HttpRequest) response.Response
	deleteMethod func(*HttpRequest) response.Response
	putMethod func(*HttpRequest) response.Response
}

func (h *handler) Get(handlerFunc func(*HttpRequest) response.Response) {
	h.getMethod = handlerFunc
}

func (h *handler) Post(handlerFunc func(*HttpRequest) response.Response) {
	h.postMethod = handlerFunc
}

func (h *handler) Delete(handlerFunc func(*HttpRequest) response.Response) {
	h.deleteMethod = handlerFunc
}

func (h *handler) Put(handlerFunc func(*HttpRequest) response.Response) {
	h.putMethod = handlerFunc
}

func (h *handler) All(handlerFunc func(*HttpRequest) response.Response) {
	h.allMethod = handlerFunc
}

func (h *handler) Request(handlerFunc func(*HttpRequest) response.Response,method []int) {
	for _,number := range method {
		if number == GET {
			h.getMethod = handlerFunc
			continue
		}

		if number == POST {
			h.postMethod = handlerFunc
			continue
		}

		if number == PUT {
			h.putMethod = handlerFunc
			continue
		}

		if number == DELETE {
			h.deleteMethod = handlerFunc
			continue
		}
	}
}