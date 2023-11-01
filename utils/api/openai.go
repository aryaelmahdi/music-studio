package api

import "github.com/sashabaranov/go-openai"

func InitAPI(APIKey string) *openai.Client {
	client := openai.NewClient(APIKey)
	return client
}
