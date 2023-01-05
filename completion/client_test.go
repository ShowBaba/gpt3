package completion_test

import (
	"fmt"
	"testing"

	"github.com/showbaba/gpt3"
	"github.com/showbaba/gpt3/completion"
)

var (
	SECRET_KEY = "sk-#######################"
	config     = gpt3.Config{
		ApiSecKey: SECRET_KEY,
	}
	client completion.Client
)

func setup() {
	client = completion.New(config)
}

func TestCompletion(t *testing.T) {
	setup()

	t.Run("Create Completion", func(t *testing.T) {
		arg := &completion.CreateCompletionReq{
			Model:            "text-davinci-003",
			Prompt:           "who is the current archbishop of Canterbury?",
			Temperature:      0.9,
			MaxTokens:        150,
			TopP:             1,
			FrequencyPenalty: 0.0,
			PresencePenalty:  0.6,
			Stop:             []string{"AI", "AI"},
		}
		resp, err := client.CreateCompletion(arg)
		if err != nil {
			t.Errorf("%v", err)
			return
		}
		fmt.Printf("response: %v", resp)
	})
}
