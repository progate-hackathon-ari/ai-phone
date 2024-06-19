package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/progate-hackathon-ari/backend/internal/repository"
)

type GenerateImageInteractor struct {
	repository repository.DataAccess
}

func NewGenerateImageInteractor(repository repository.DataAccess) *GenerateImageInteractor {
	return &GenerateImageInteractor{
		repository: repository,
	}
}

func (i *GenerateImageInteractor) GenerateImage(ctx context.Context, roomId string, connectionId string, prompt string) (string, error) {
	status := "success"
	Url, err := GetImageUrl(ctx, roomId, connectionId, prompt)
	if err != nil {
		status = "failed"
		return nil, errors.Join(err, fmt.Errorf("failed to generate image"))
	}

	fmt.Printf("Url: %v\n", Url)

	return status, nil
}

func GetImageUrl(ctx context.Context, roomId string, connectionId string, prompt string) (string, error) {
	//TODO：bedrockAPI叩く

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           "di",
		SharedConfigState: session.SharedConfigEnable,
	}))

	uploader := s3manager.NewUploader(sess)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("image"),
		Key:    aws.String(`image/` + roomId + `/` + connectionId + `/` + prompt + `.png`),
		Body:   file,
	})

	if err != nil {
		return nil, errors.Join(err, fmt.Errorf("failed to upload image"))
	}

	return `/image/` + roomId + `/` + connectionId + `/` + prompt + `.png`, err
}

func sendUrl(ctx context.Context, url string) (string, error) {

	return nil, nil
}
