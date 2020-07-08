package response

import "net/http"

type Response interface {
	SendResponse(http.ResponseWriter) error
	WriteHeader(http.ResponseWriter)
	SetHeader(map[string]string) jsonResponse
	AddHeader(string,string) jsonResponse
	GetStatusCode() int8
	SetStatusCode(int8) jsonResponse
}