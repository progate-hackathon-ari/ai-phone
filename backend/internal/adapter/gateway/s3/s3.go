package s3

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/progate-hackathon-ari/backend/cmd/config"
)

type S3 interface {
	UplaodImage(ctx context.Context, filename string, image []byte) error
}

type S3Repo struct {
	client *s3.Client
}

func NewS3Repo(config aws.Config) *S3Repo {
	return &S3Repo{
		client: s3.NewFromConfig(config),
	}
}

type AWSLess struct{}

func NewAWSLess() *AWSLess {
	return &AWSLess{}
}

func (r *AWSLess) UplaodImage(ctx context.Context, filename string, image []byte) error {
	return nil
}

func (r *S3Repo) UplaodImage(ctx context.Context, filename string, image []byte) error {
	_, err := r.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(config.Config.Aws.S3BucketName),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(image),
	})
	if err != nil {
		return err
	}
	return nil
}
