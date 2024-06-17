package container

import (
	"github.com/progate-hackathon-ari/backend/internal/driver/db"
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
		panic(err)
	}

	return r
}
