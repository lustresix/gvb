package gpt_api

import (
	"errors"
	"fmt"
	"gbv2/config/gpt"
	"gbv2/models/res"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"io"
)

type ChatReq struct {
	Content string `json:"content"`
}

func (GptApi) Chat(c *gin.Context) {
	var cq ChatReq
	err := c.ShouldBindJSON(&cq)
	if err != nil {
		res.FailWithCode(res.ErrorParameterTransfer, c)
		return
	}

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 20,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: cq.Content,
			},
		},
		Stream: true,
	}

	stream, err := gpt.GPT.CreateChatCompletionStream(c, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	fmt.Printf("Stream response: ")

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return
		}
		res.OKWithMsg(response.Choices[0].Delta.Content, c)
		fmt.Printf(response.Choices[0].Delta.Content)
	}

}
