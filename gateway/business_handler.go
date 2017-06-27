package gateway

import (
	"bytes"
	"net/http"
	"io/ioutil"
	"errors"
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
