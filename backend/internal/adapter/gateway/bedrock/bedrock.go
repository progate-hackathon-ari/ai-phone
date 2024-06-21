package bedrock

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

type Bedrock interface {
	// SDXL1.0
	GenerateImageFromText(ctx context.Context, prompt, negativePrompt, style string) ([][]byte, error)
	// Cloude
	BuildPrompt(ctx context.Context, prompt string) (*Prompts, error)
}

type BedrockService struct {
	brc *bedrockruntime.Client
}

func NewBedRock(config aws.Config) *BedrockService {
	return &BedrockService{
		brc: bedrockruntime.NewFromConfig(config),
	}
}

type AILessService struct {
}

func NewAILessBedRock() *AILessService {
	return &AILessService{}
}

func (r *AILessService) GenerateImageFromText(ctx context.Context, prompt, negativePrompt, style string) ([][]byte, error) {
	return nil, nil
}

// Cloude
func (r *AILessService) BuildPrompt(ctx context.Context, prompt string) (*Prompts, error) {
	return nil, nil
}
