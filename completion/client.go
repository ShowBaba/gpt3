package completion

import (
	"encoding/json"

	"github.com/showbaba/gpt3"
)

type Client interface {
	CreateCompletion(arg *CreateCompletionReq) (res CreateCompletionResponse, err error)
}

type apiImpl gpt3.Context

func New(g gpt3.Config) Client {
	ctx := apiImpl(gpt3.NewContext(g))
	return &ctx
}

func (c *apiImpl) CreateCompletion(arg *CreateCompletionReq) (res CreateCompletionResponse, err error) {
	reqData := gpt3.Request{
		Method: "POST",
		Route:  []string{"completion", "create"},
	}
	resp, err := (*gpt3.Context)(c).SendRequest(reqData, arg)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &res)
	return
}
