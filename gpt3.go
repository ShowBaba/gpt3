package gpt3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	BASE_URL = "https://api.openai.com/v1"
)

type Config struct {
	ApiSecKey string
}

type Context struct {
	Config Config
}

func NewContext(c Config) Context {
	return Context{c}
}

var Endpoints = map[string]map[string]string{
	"completion": {
		"create": "/completions",
	},
}

func (ctx *Context) GetEndpoint(endpoint []string) string {
	return Endpoints[endpoint[0]][endpoint[1]]
}

type Request struct {
	Method string
	Route  []string
}

func (ctx *Context) GenerateURL(route []string) string {
	endpoint := ctx.GetEndpoint(route)
	return fmt.Sprintf(`%s%s`, BASE_URL, endpoint)
}

func (ctx *Context) SendRequest(reqData Request, data interface{}) (b []byte, err error) {
	url := ctx.GenerateURL(reqData.Route)
	var (
		req  *http.Request
		body *bytes.Buffer
	)
	if data != nil {
		jsonBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(jsonBytes)
		req, err = http.NewRequest(reqData.Method, url, body)
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest(reqData.Method, url, nil)
		if err != nil {
			return nil, err
		}
	}
	req.Header.Add("Authorization", fmt.Sprintf(`Bearer %s`, ctx.Config.ApiSecKey))
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	b, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusOK {
		return b, nil
	}
	return b, APIError{
		Content: string(b),
	}
}

type APIError struct {
	Content string
}

func (e APIError) Error() string {
	return fmt.Sprintf(`unexpected error occured; error: %v;`, e.Content)
}

