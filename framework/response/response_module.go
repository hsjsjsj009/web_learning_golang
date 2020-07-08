package response

import "net/http"

type Response interface {
	SendResponse(http.ResponseWriter) error
	WriteHeader(http.ResponseWriter)
	SetHeader(map[string]string)
	AddHeader(string,string) jsonResponse
}