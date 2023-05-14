package chat

import (
	"context"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

type Bot struct{}

func New() *Bot {
	return &Bot{}
}

func (*Bot) Message(prompt string) (string, error) {
	client := openai.NewClient(os.Getenv("API_OPENAI"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil

}

func (*Bot) MessageRefferSystem(system, prompt string) (string, error) {
	client := openai.NewClient(os.Getenv("API_OPENAI"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: system,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil

}
