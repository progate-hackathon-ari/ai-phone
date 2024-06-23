package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/progate-hackathon-ari/backend/internal/adapter/gateway/bedrock"
	"github.com/progate-hackathon-ari/backend/internal/entities/model"
	"github.com/progate-hackathon-ari/backend/pkg/log"
)

type AnswerResponse struct {
	IsAllUserAnswered bool `json:"is_all_user_answered"`
}

func (i *GameInteractor) ImageGenerate(ctx context.Context, roomID, prompt string) error {
	// 2回目も防ぐ
	if Rooms[roomID].Players[i.client.info.ConnectionID].IsAnswered {
		return fmt.Errorf("already answered")
	}

	room, err := i.repo.GetRoom(ctx, roomID)
	if err != nil {
		return err
	}

	if err := i.repo.CreateIngamePrompt(ctx, &model.InGamePrompt{
		RoomID:       roomID,
		ConnectionID: i.client.info.ConnectionID,
		GameIndex:    room.CurrentGame,
		Prompt:       prompt,
	}); err != nil {
		return err
	}

	var resultPrompt *bedrock.Prompts
	var images [][]byte
	var retryErr error

	for r := range 5 {
		resultPrompt, retryErr = i.bedrock.BuildPrompt(ctx, strings.Join([]string{
			prompt,
			// 暗黙的な内部の追加プロンプトはここに書く
		}, ","))
		if retryErr != nil {
			log.Error(ctx, "failed to build prompt", retryErr)
			continue
		}

		log.Info(ctx, "resultPrompt", "prompt", resultPrompt.Prompt, "negative", resultPrompt.NegativePrompt)

		images, retryErr = i.bedrock.GenerateImageFromText(ctx, resultPrompt.Prompt, resultPrompt.NegativePrompt, room.ExtraPrompt)
		if retryErr == nil {
			break
		}

		log.Error(ctx, "failed to generate image", retryErr, "retry", r)
		err = retryErr
	}

	if err != nil {
		return err
	}

	if err := i.s3.UplaodImage(ctx, fmt.Sprintf("%s/%s/%d.jpg", roomID, i.client.info.ConnectionID, room.CurrentGame), images[0]); err != nil {
		return err
	}

	Rooms[roomID].Players[i.client.info.ConnectionID].IsAnswered = true

	data, err := json.Marshal(&AnswerResponse{
		IsAllUserAnswered: IsAnswered(roomID),
	})
	if err != nil {
		return err
	}

	return BroadcastInRoom(roomID, data)
}
