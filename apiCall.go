package main

import (
	"context"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func sendMessage(msg string) (string, error) {
	token := os.Getenv("OPEN_AI_SECRET_KEY")
	client := openai.NewClient(token)
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: msg,
			},
		},
	}

	resp, err := client.CreateChatCompletion(context.Background(), req)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
	// return res, nil
}
