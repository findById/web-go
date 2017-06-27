package message

import (
	"net/http"
	"web-go/common"
	"encoding/json"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	result := common.Response{}
	result.StatusCode = "200"
	result.Message = "success"
	result.Result = ""
	json.NewEncoder(w).Encode(result)
}
