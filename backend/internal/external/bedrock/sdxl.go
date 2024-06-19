package bedrock

import (
	"context"
	"encoding/base64"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

const SDXLModelID = "stability.stable-diffusion-xl-v1"

type Prompt struct {
	Text string `json:"text"`
}

type ReqBody struct {
	TextPrompts []Prompt `json:"text_prompts"`
	CfgScale    int      `json:"cfg_scale"`
	Seed        int      `json:"seed"`
	Steps       int      `json:"steps"`
	Width       int      `json:"width"`
	Height      int      `json:"height"`
}

type Artifact struct {
	Seed   string `json:"seed"`
	Base64 string `json:"base64"`
}

type Result struct {
	Result    string     `json:"result"`
	Artifacts []Artifact `json:"artifacts"`
}

func (s *BedrockService) GenerateImageFromText(ctx context.Context, prompt string) ([][]byte, error) {
	body := ReqBody{
		TextPrompts: []Prompt{
			{
				Text: prompt,
			},
		},
		// TODO: 調整
		CfgScale: 30,
		Seed:     749500536,
		Steps:    130,
		Width:    512,
		Height:   512,
	}

	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	resp, err := s.brc.InvokeModel(ctx, &bedrockruntime.InvokeModelInput{
		ModelId:     aws.String(`stability.stable-diffusion-xl-v1`),
		ContentType: aws.String(`application/json`),
		Accept:      aws.String(`application/json`),
		Body:        reqBody,
	})

	if err != nil {
		return nil, err
	}

	response := Result{}
	_ = json.Unmarshal(resp.Body, &response)

	if response.Result != `success` {
		panic(response.Result)

	}

	var result [][]byte

	for _, artifact := range response.Artifacts {
		decodedData, err := base64.StdEncoding.DecodeString(artifact.Base64)
		if err != nil {
			return nil, err
		}
		result = append(result, decodedData)
	}

	return result, nil
}
