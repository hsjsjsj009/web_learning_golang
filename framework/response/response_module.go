package response

import "net/http"

type Response interface {
	WriteResponse(http.ResponseWriter) ([]byte,error)
	WriteHeader(http.ResponseWriter)
	SetHeader(map[string]string) Response
	AddHeader(string,string) Response
	GetStatusCode() int
	SetStatusCode(int) Response
}