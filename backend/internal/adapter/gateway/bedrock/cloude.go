package bedrock

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"

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

const DefaultRetry int = 5

type Prompts struct {
	Text           string   `json:"text"`
	Prompt         []string `json:"prompts"`
	NegativePrompt []string `json:"negative_prompts"`
}

func (s *BedrockService) BuildPrompt(ctx context.Context, prompt string) (*Prompts, error) {
	for range DefaultRetry {
		prompt, err := s.requestPrompt(ctx, fmt.Sprintf(`Human: <text> %s</text>
		I want to translate it into English, generate prompts and negative prompts in JSON format, output only the data, and recreate {{text}} with stable diffusion.
		
		Assistant:
		`, prompt))
		if err != nil {
			return nil, err
		}

		reg := regexp.MustCompile("```json([^`]*)```")
		matches := reg.FindAllStringSubmatch(prompt, -1)

		if len(matches) == 0 {
			continue
		}

		jsonPrompt := strings.ReplaceAll(strings.ReplaceAll(matches[0][1], "```json", ""), "```", "")
		log.Println(jsonPrompt)

		var prompts Prompts
		err = json.Unmarshal([]byte(jsonPrompt), &prompts)
		if err == nil {
			prompts.Prompt = append(prompts.Prompt, prompts.Text)
			return &prompts, nil
		}
	}
	return nil, fmt.Errorf("failed to build prompt")
}
