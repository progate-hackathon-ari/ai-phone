package bedrock

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types"
)

type ClaudeRequest struct {
	Prompt            string   `json:"prompt"`
	MaxTokensToSample int      `json:"max_tokens_to_sample"`
	Temperature       float64  `json:"temperature,omitempty"`
	StopSequences     []string `json:"stop_sequences,omitempty"`
}

type Content struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type Message struct {
	Role    string    `json:"role"`
	Content []Content `json:"content"`
}

type RequestBodyClaude3 struct {
	MaxTokensToSample int       `json:"max_tokens"`
	Temperature       float64   `json:"temperature,omitempty"`
	AnthropicVersion  string    `json:"anthropic_version"`
	Messages          []Message `json:"messages"`
}

type Delta struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type ResponseClaude3 struct {
	Type  string `json:"type"`
	Index int    `json:"index"`
	Delta Delta  `json:"delta"`
}

func (s *BedrockService) requestPrompt(ctx context.Context, prompt string) (string, error) {
	payload := RequestBodyClaude3{
		MaxTokensToSample: 2048,
		AnthropicVersion:  "bedrock-2023-05-31",
		Temperature:       0.9,
		Messages: []Message{
			{
				Role: "user",
				Content: []Content{
					{
						Type: "text",
						Text: prompt,
					},
				},
			},
		},
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	output, err := s.brc.InvokeModelWithResponseStream(
		ctx,
		&bedrockruntime.InvokeModelWithResponseStreamInput{
			Body:        payloadBytes,
			ModelId:     aws.String("anthropic.claude-3-haiku-20240307-v1:0"),
			ContentType: aws.String("application/json"),
			Accept:      aws.String("application/json"),
		},
	)

	if err != nil {
		return "", err
	}

	defer output.GetStream().Close()

	var response []string

	for event := range output.GetStream().Events() {
		switch v := event.(type) {
		case *types.ResponseStreamMemberChunk:
			var resp ResponseClaude3
			err := json.NewDecoder(bytes.NewReader(v.Value.Bytes)).Decode(&resp)
			if err != nil {
				return "", err
			}

			response = append(response, resp.Delta.Text)
		case *types.UnknownUnionMember:
			return "", err

		default:
			return "", err
		}
	}

	return strings.Join(response, ""), nil
}

const DefaultRetry int = 5

type Prompts struct {
	Prompt         string `json:"prompt"`
	NegativePrompt string `json:"negative_prompt"`
}

func (s *BedrockService) BuildPrompt(ctx context.Context, prompt string) (*Prompts, error) {
	for range DefaultRetry {
		prompt, err := s.requestPrompt(ctx, fmt.Sprintf(`<text> %s</text>
I want to reproduce {{text}} in Stable diffusion by generating prompts and negative prompts in JSON format and outputting data only
`, prompt))
		if err != nil {
			return nil, err
		}

		reg := regexp.MustCompile(`{[^{}]*}`)
		matches := reg.FindAllStringSubmatch(prompt, -1)

		if len(matches) == 0 {
			continue
		}

		fmt.Println(matches[0][0])

		var prompts Prompts
		err = json.Unmarshal([]byte(matches[0][0]), &prompts)
		if err == nil {
			return &prompts, nil
		}

		log.Println(err)
	}
	return nil, fmt.Errorf("failed to build prompt")
}
