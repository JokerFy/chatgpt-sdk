package v1

import (
	"encoding/json"
	"fmt"
)

const (
	Embeddings = "https://api.openai.com/v1/embeddings"
)

type EmbeddingsRequest struct {
	Model string `json:"model"`          // ID of the model to use. You can use the List models API to see all of your available models, or see our Model overview for descriptions of them.
	Input string `json:"input"`          // Input text to get embeddings for, encoded as a string or array of tokens. To get embeddings for multiple inputs in a single request, pass an array of strings or array of token arrays. Each input must not exceed 8192 tokens in length.
	User  string `json:"user,omitempty"` // A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse
}

type EmbeddingsResponse struct {
	Object string `json:"object"`
	Data   []struct {
		Object    string    `json:"object"`
		Embedding []float64 `json:"embedding"`
		Index     int       `json:"index"`
	} `json:"data"`
	Model string `json:"model"`
	Usage struct {
		PromptTokens int `json:"prompt_tokens"`
		TotalTokens  int `json:"total_tokens"`
	} `json:"usage"`
}

// Embeddings Creates an embedding vector representing the input text.
func (chat *ChatGpt) Embeddings(req EmbeddingsRequest) (response EmbeddingsResponse, err error) {
	resp, err := chat.Post(Embeddings, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(resp, &response)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
