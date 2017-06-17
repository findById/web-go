package message

import (
	"net/http"
	"web-go/common"
	"encoding/json"
	"io/ioutil"
)

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	result := common.Response{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		result.StatusCode = "201"
		result.Message = err.Error()
		json.NewEncoder(w).Encode(result)
		return
	}

	result.StatusCode = "200"
	result.Message = "success"
	result.Result = string(body)
	json.NewEncoder(w).Encode(result)
}
