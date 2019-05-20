package server

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

const baseUrl = "https://api.netease.im/nimserver/"

type config struct {
	appKey    string
	appSecret string
}

type client struct {
	config *config
}

var flag = false
var instance *client

func ClientInstance() (*client, error) {
	if flag {
		return instance, nil
	}
	return nil, errors.New("错误")
}

func NewClient(appKey, appSecret string) *client {
	flag = true
	instance = &client{
		config: &config{
			appKey:    appKey,
			appSecret: appSecret,
		},
	}
	return instance
}

func (client *client) Do(request Request) (response string, err error) {
	val, err := request.GetQuery()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	body := bytes.NewBuffer([]byte(val.Encode()))
	httpRequest, err := http.NewRequest(request.GetMethod(), request.GetUrl(baseUrl), body)
	if err != nil {
		return "", err
	}
	headers := request.GetHeader(client.config.appKey, client.config.appSecret)
	for key, value := range headers {
		httpRequest.Header.Set(key, value)
	}
	httpClient := &http.Client{}
	resp, err := httpClient.Do(httpRequest)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}
	content, err := ioutil.ReadAll(resp.Body)
	return string(content), err
}
