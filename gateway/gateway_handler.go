package gateway

import (
	"web-go/common"
	"encoding/json"
	"net/http"
	"bytes"
	"strings"
	"errors"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	result := common.Response{}

	body := new(Request)
	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		result.StatusCode = "201"
		result.Message = err.Error()
		json.NewEncoder(w).Encode(result)
		return
	}
	method, err := parseMethod(body.Method)
	if err != nil {
		result.StatusCode = "202"
		result.Message = err.Error()
		json.NewEncoder(w).Encode(result)
		return
	}

	bizContent, err := parseBizContent(body.BizContent)
	if err != nil {
		result.StatusCode = "202"
		result.Message = "'Illegal response data."
		json.NewEncoder(w).Encode(result)
		return
	}

	bizResult, err := Post("http://localhost:8080/"+method, bizContent)
	if err != nil {
		result.StatusCode = "502"
		result.Message = "remote access failed"
		json.NewEncoder(w).Encode(result)
		return
	}
	err = json.Unmarshal(bizResult, &result.Result)
	if err != nil {
		result.StatusCode = "502"
		result.Message = "remote access failed"
		json.NewEncoder(w).Encode(result)
		return
	}
	result.StatusCode = "200"
	result.Message = "success"
	json.NewEncoder(w).Encode(result)
}

func parseBizContent(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func parseMethod(method string) (string, error) {
	if len(method) <= 0 {
		return "", errors.New("'method' parsing failed.")
	}
	return strings.Replace(method, ".", "/", -1), nil
}
