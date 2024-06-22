package ariaws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/progate-hackathon-ari/backend/pkg/log"
)

func NewConfig() aws.Config {
	sdkConfig, err := config.LoadDefaultConfig(context.Background(), config.WithRegion("us-west-2"))
	if err != nil {
		log.Fatal(context.Background(), err.Error())
	}

	return sdkConfig
}
