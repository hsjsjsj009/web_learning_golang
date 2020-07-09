package response

import (
	"net/http"
)

type redirect struct {
	url string
	statusCode int
	req *http.Request
}

func (r redirect) WriteResponse(writer http.ResponseWriter) ([]byte, error) {
	statusCode := r.statusCode
	if statusCode == 0 {
		statusCode = http.StatusFound
	}
	http.Redirect(writer,r.req,r.url,statusCode)
	return nil,nil
}

func (r redirect) WriteHeader(writer http.ResponseWriter) {
}

func (r redirect) SetHeader(m map[string]string) Response {
	return r
}

func (r redirect) AddHeader(key string, value string) Response {
	return r
}

func (r redirect) GetStatusCode() int {
	return r.statusCode
}

func (r redirect) SetStatusCode(i int) Response {
	r.statusCode = i
	return r
}

func Redirect(url string,request *http.Request) redirect {
	return redirect{url: url,req: request}
}


