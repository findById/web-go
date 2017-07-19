package gateway

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

func Post(url string, data []byte) ([]byte, error) {
	bizResponse, err := http.Post(url, "application/json", bytes.NewReader(data))
	defer bizResponse.Body.Close()
	if err != nil {
		return nil, err
	}
	if bizResponse.StatusCode != http.StatusOK {
		return nil, errors.New("remote status error")
	}
	bizResult, err := ioutil.ReadAll(bizResponse.Body)
	if err != nil {
		return nil, err
	}
	return bizResult, nil
}
