package response

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	header map[string]string
	Data map[string]interface{} `json:"data"`
	statusCode int8
}

func JsonResponse(data map[string]interface{}) jsonResponse{
	return jsonResponse{Data: data}
}

func (j jsonResponse) SetHeader(m map[string]string) jsonResponse {
	j.header = m
	return j
}

func (j jsonResponse) GetStatusCode() int8 {
	return j.statusCode
}

func (j jsonResponse) SetStatusCode(code int8) jsonResponse{
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

func (j jsonResponse) AddHeader(key string,value string) jsonResponse {
	j.header[key] = value
	return j
}

func (j jsonResponse) SendResponse(w http.ResponseWriter) error {
	js,err := json.Marshal(j)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)

	return err
}
