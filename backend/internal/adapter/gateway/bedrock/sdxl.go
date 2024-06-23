package bedrock

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"math/rand"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

const SDXLModelID = "stability.stable-diffusion-xl-v1"

type StylePreset string

const (
	Style3DModel    = StylePreset("3d-model")
	StyleAnalogFilm = StylePreset("analog-film")
	StyleAnime      = StylePreset("anime")
	StyleCinematic  = StylePreset("cinematic")
	StyleComicBook  = StylePreset("comic-book")

	StyleDigitalArt = StylePreset("digital-art")
	StyleEnhance    = StylePreset("enhance")
	StyleFantasyArt = StylePreset("fantasy-art")
	StyleIsometric  = StylePreset("isometric")
	StyleLineArt    = StylePreset("line-art")

	StyleLowPoly          = StylePreset("low-poly")
	StyleModelingCompound = StylePreset("modeling-compound")
	StyleNeonPunk         = StylePreset("neon-punk")
	StyleOrigami          = StylePreset("origami")
	StylePhotographic     = StylePreset("photographic")

	StylePixelArt = StylePreset("pixel-art")

	StyleTileTexture = StylePreset("tile-texture")
)

var Styles = map[int]StylePreset{
	0:  Style3DModel,
	1:  StyleAnalogFilm,
	2:  StyleAnime,
	3:  StyleCinematic,
	4:  StyleComicBook,
	5:  StyleDigitalArt,
	6:  StyleEnhance,
	7:  StyleFantasyArt,
	8:  StyleIsometric,
	9:  StyleLineArt,
	10: StyleLowPoly,
	11: StyleModelingCompound,
	12: StyleNeonPunk,
	13: StyleOrigami,
	14: StylePhotographic,
	15: StylePixelArt,
	16: StyleTileTexture,
	17: Style3DModel,
}

func RandomStyle() StylePreset {
	return Styles[rand.Intn(len(Styles))]
}

type Prompt struct {
	Text   string  `json:"text"`
	Weight float32 `json:"weight"`
}

type ReqBody struct {
	TextPrompts []Prompt    `json:"text_prompts"`
	CfgScale    int         `json:"cfg_scale"`
	Seed        int         `json:"seed"`
	Steps       int         `json:"steps"`
	Width       int         `json:"width"`
	Height      int         `json:"height"`
	StylePreset StylePreset `json:"style_preset"`
}

type Artifact struct {
	Seed   string `json:"seed"`
	Base64 string `json:"base64"`
}

type Result struct {
	Result    string     `json:"result"`
	Artifacts []Artifact `json:"artifacts"`
}

func (s *BedrockService) GenerateImageFromText(ctx context.Context, prompt, negativePrompt, style string) ([][]byte, error) {
	if style == "\n" {
		style = string(RandomStyle())
	}

	body := ReqBody{
		TextPrompts: []Prompt{
			{
				Text:   prompt,
				Weight: 0.7,
			},
			{
				Text:   negativePrompt,
				Weight: -0.7,
			},
		},
		// TODO: 調整
		CfgScale:    5,
		Seed:        0,
		Steps:       50,
		Width:       1216,
		Height:      832,
		StylePreset: StylePreset(style),
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
