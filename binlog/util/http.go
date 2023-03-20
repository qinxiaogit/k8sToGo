package util

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

type Http struct {
}

//https://api.weixin.qq.com/cgi-bin/token

func (h *Http) Get(url string, params map[string]string) (string, error) {
	client := resty.New()

	resp, err := client.R().SetQueryParams(params).EnableTrace().
		Get(url)

	if err != nil {
		return "", err
	}

	return string(resp.Body()), err
}

func (h *Http)PostJson(url string, params map[string]interface{}) (string, error) {
	client := resty.New()

	jsonObj, _ := json.Marshal(params)

	resp, err := client.R().SetHeader("Content-Type", "application/json").SetBody(jsonObj).EnableTrace().Post(url)

	if err != nil {
		return "", err
	}

	return string(resp.Body()), err
}
