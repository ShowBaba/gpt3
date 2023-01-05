# GPT-3 SDK

The goal of this project is to provide a simple and easy-to-use tool for developers who want to access the GPT-3 API from their Go projects.

I have focused on the completions endpoint for this initial release of the SDK, but I plan to add support for additional GPT-3 endpoints in the future.

I welcome contributions from the community to help improve and expand the project. The repository includes a detailed readme file with instructions on how to use the SDK and how to contribute to the project.

## Installation
```sh
go get github.com/ShowBaba/gpt3
```

This SDK is built so you can import relavant namespace(s) only.

## Exported sub-packages
* `go get github.com/ShowBaba/gpt3/completion`

With the base at:
* `go get github.com/ShowBaba/gpt3`

## Usage 

### Configuration
```go
  import "github.com/ShowBaba/gpt3"

  func main() {
    config  := gpt3.Config{
      ApiSecKey: "sk-#######################",
    }
  }
```

### Making API calls

## Completion

### `Create Completion`
This describes how to create a completion for the provided prompt and parameters

```go
  import "github.com/showbaba/gpt3/completion"

  client := completion.New(config)

  arg := &completion.CreateCompletionReq{
    Model:            "text-davinci-003",
    Prompt:           "who is the current archbishop of Canterbury",
    Temperature:      0.9,
    MaxTokens:        150,
    TopP:             1,
    FrequencyPenalty: 0.0,
    PresencePenalty:  0.6,
    Stop:             []string{"AI", "AI"},
  }
  resp, err := client.CreateCompletion(arg)
  if err != nil {
    t.Errorf("unexpected error occured; err: %v", err)
    return
  }
  fmt.Printf("response: %v", resp)
```