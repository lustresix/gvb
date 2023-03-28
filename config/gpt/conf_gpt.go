package gpt

import (
	"context"
	"gbv2/config/log"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

var GPT *openai.Client

func ConnectGpt() {
	client := openai.NewClient(viper.GetString("gpt.key"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo0301,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		log.Errorw("GPT连接失败", "err", err)
		return
	}

	GPT = client
	log.Infow("GPT连接成功" + resp.Choices[0].Message.Content)
}
