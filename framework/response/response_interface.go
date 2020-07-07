package response

import "net/http"

type Response interface {
	SendResponse(*http.ResponseWriter)
}