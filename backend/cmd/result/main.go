// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// snippet-start:[gov2.bedrock-runtime.Hello]

package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/progate-hackathon-ari/backend/internal/external/bedrock"
	"github.com/progate-hackathon-ari/backend/internal/external/s3"
)

// Each model provider defines their own individual request and response formats.
// For the format, ranges, and default values for the different models, refer to:
// https://docs.aws.amazon.com/bedrock/latest/userguide/model-parameters.html

type ClaudeRequest struct {
	Prompt            string `json:"prompt"`
	MaxTokensToSample int    `json:"max_tokens_to_sample"`
	// Omitting optional request parameters
}

type ClaudeResponse struct {
	Completion string `json:"completion"`
}

// main uses the AWS SDK for Go (v2) to create an Amazon Bedrock Runtime client
// and invokes Anthropic Claude 2 inside your account and the chosen region.
// This example uses the default settings specified in your shared credentials
// and config files.
func main() {

	region := flag.String("region", "us-east-1", "The AWS region")
	flag.Parse()

	fmt.Printf("Using AWS region: %s\n", *region)

	sdkConfig, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(*region))
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}
	brs := bedrock.NewBedRock(sdkConfig)
	prompt, err := brs.BuildPrompt(context.Background(), "おいしいラーメンの画像、ビールを添えて, 1000文字以内で, 8bit風に")
	if err != nil {
		panic(err)
	}

	result, err := brs.GenerateImageFromText(context.Background(), prompt)
	if err != nil {
		panic(err)
	}

	s3bs := s3.NewS3Repo(sdkConfig)

	for _, i := range result {
		if err := s3bs.UplaodImage(context.Background(), "example/"+time.Now().Format(time.RFC3339)+".jpg", i); err != nil {
			panic(err)
		}
	}
}
