package api

import (
	"encoding/json"
	"net/http"
)

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	result := Response{}
	result.StatusCode = "200"
	result.Message = "success"
	result.Result = ""
	json.NewEncoder(w).Encode(result)
}
