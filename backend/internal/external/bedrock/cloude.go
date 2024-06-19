package bedrock

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

type ClaudeRequest struct {
	Prompt            string   `json:"prompt"`
	MaxTokensToSample int      `json:"max_tokens_to_sample"`
	Temperature       float64  `json:"temperature,omitempty"`
	StopSequences     []string `json:"stop_sequences,omitempty"`
}

type ClaudeResponse struct {
	Completion string `json:"completion"`
}

func (s *BedrockService) requestPrompt(ctx context.Context, prompt string) (string, error) {
	modelId := "anthropic.claude-v2"

	body, err := json.Marshal(ClaudeRequest{
		Prompt:            prompt,
		MaxTokensToSample: 200,
		Temperature:       0.5,
		StopSequences:     []string{"\n\nHuman:"},
	})

	if err != nil {
		log.Fatal("failed to marshal", err)
	}

	output, err := s.brc.InvokeModel(ctx, &bedrockruntime.InvokeModelInput{
		ModelId:     aws.String(modelId),
		ContentType: aws.String("application/json"),
		Accept:      aws.String("application/json"),
		Body:        body,
	})

	if err != nil {
		return "", err
	}

	var response ClaudeResponse
	if err := json.Unmarshal(output.Body, &response); err != nil {
		log.Fatal("failed to unmarshal", err)
	}

	return response.Completion, nil
}

func (s *BedrockService) TranslateToEnglish(ctx context.Context, prompt string) (string, error) {
	return s.requestPrompt(ctx, fmt.Sprintf("Human: Please translate this to english: %s \n\nAssistant:", prompt))
}

func (s *BedrockService) BuildPrompt(ctx context.Context, prompt string) (string, error) {
	return s.requestPrompt(ctx, fmt.Sprintf("Human: Please output the best prompt in English that creates an image associated with this string using stable deffusion.: %s \n\nAssistant:", prompt))
}
