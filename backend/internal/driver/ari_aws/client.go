package ariaws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

type AWSConnect struct {
	accountID       string
	endpoint        string
	accessKeyID     string
	accessKeySecret string
}

func New(
	accountID,
	endpoint,
	accessKeyID,
	accessKeySecret string,
) *AWSConnect {
	return &AWSConnect{
		accountID:       accountID,
		endpoint:        endpoint,
		accessKeyID:     accessKeyID,
		accessKeySecret: accessKeySecret,
	}
}

func (a *AWSConnect) Config(ctx context.Context) (aws.Config, error) {
	return config.LoadDefaultConfig(
		ctx,
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				a.accessKeyID,
				a.accessKeySecret,
				"",
			)),
		config.WithRegion("auto"),
	)
}
