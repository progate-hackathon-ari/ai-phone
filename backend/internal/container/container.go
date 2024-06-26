package container

import (
	"log"

	"github.com/progate-hackathon-ari/backend/internal/adapter/gateway/bedrock"
	"github.com/progate-hackathon-ari/backend/internal/adapter/gateway/repository"
	"github.com/progate-hackathon-ari/backend/internal/adapter/gateway/repository/gorm"
	"github.com/progate-hackathon-ari/backend/internal/adapter/gateway/s3"
	ariaws "github.com/progate-hackathon-ari/backend/internal/driver/ari_aws"
	"github.com/progate-hackathon-ari/backend/internal/driver/db"
	"github.com/progate-hackathon-ari/backend/internal/usecase"
	"go.uber.org/dig"
)

var container *dig.Container

type provideArg struct {
	constructor any
	opts        []dig.ProvideOption
}

func NewContainer() error {
	container = dig.New()

	args := []provideArg{
		{constructor: db.Connect, opts: []dig.ProvideOption{}},
		{constructor: db.NewGORM, opts: []dig.ProvideOption{}},
		{constructor: gorm.NewGormDB, opts: []dig.ProvideOption{dig.As(new(repository.DataAccess))}},
		{constructor: ariaws.NewConfig, opts: []dig.ProvideOption{}},
		{constructor: s3.NewS3Repo, opts: []dig.ProvideOption{dig.As(new(s3.S3))}},
		{constructor: bedrock.NewBedRock, opts: []dig.ProvideOption{dig.As(new(bedrock.Bedrock))}},
		{constructor: usecase.NewRoomInteractor, opts: []dig.ProvideOption{}},
	}

	for _, arg := range args {
		if err := container.Provide(arg.constructor, arg.opts...); err != nil {
			return err
		}
	}

	return nil
}

func NewAILessContainer() error {
	container = dig.New()

	args := []provideArg{
		{constructor: db.Connect, opts: []dig.ProvideOption{}},
		{constructor: db.NewGORM, opts: []dig.ProvideOption{}},
		{constructor: gorm.NewGormDB, opts: []dig.ProvideOption{dig.As(new(repository.DataAccess))}},
		{constructor: s3.NewAWSLess, opts: []dig.ProvideOption{dig.As(new(s3.S3))}},
		{constructor: bedrock.NewAILessBedRock, opts: []dig.ProvideOption{dig.As(new(bedrock.Bedrock))}},
		{constructor: usecase.NewRoomInteractor, opts: []dig.ProvideOption{}},
	}

	for _, arg := range args {
		if err := container.Provide(arg.constructor, arg.opts...); err != nil {
			return err
		}
	}

	return nil
}

func Invoke[T any]() T {
	var r T
	if err := container.Invoke(func(t T) error {
		r = t
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	return r
}
