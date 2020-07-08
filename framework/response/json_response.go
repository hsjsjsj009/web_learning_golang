package response

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	header map[string]string
	data map[string]interface{}
}

func (j jsonResponse) SetHeader(m map[string]string) {
	panic("implement me")
}

func (j jsonResponse) WriteHeader(writer http.ResponseWriter) {
	if j.header != nil {
		for key,value := range j.header{
			writer.Header().Set(key,value)
		}
	}
}

func NewJsonResponse(data map[string]interface{}) jsonResponse {
	return jsonResponse{data: data}
}

func (j jsonResponse) AddHeader(key string,value string) jsonResponse {
	j.header[key] = value
	return j
}

func (j jsonResponse) SendResponse(w http.ResponseWriter) error {
	js,err := json.Marshal(j.data)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)

	return err
}
