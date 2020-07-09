package response

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	header map[string]string
	Data map[string]interface{} `json:"data"`
	statusCode int
}

func JsonResponse(data map[string]interface{}) jsonResponse{
	return jsonResponse{Data: data}
}

func (j jsonResponse) SetHeader(m map[string]string) Response {
	j.header = m
	return j
}

func (j jsonResponse) GetStatusCode() int {
	return j.statusCode
}

func (j jsonResponse) SetStatusCode(code int) Response {
	j.statusCode = code
	return j
}

func (j jsonResponse) WriteHeader(writer http.ResponseWriter) {
	if j.header != nil {
		for key,value := range j.header {
			writer.Header().Set(key,value)
		}
	}
}

func (j jsonResponse) AddHeader(key string,value string) Response {
	if j.header == nil {
		j.header = map[string]string{}
	}
	j.header[key] = value
	return j
}

func (j jsonResponse) WriteResponse(w http.ResponseWriter) ([]byte,error) {
	js,err := json.Marshal(j)

	if err != nil {
		return nil,err
	}

	w.Header().Set("Content-Type", "application/json")

	return js,nil
}
