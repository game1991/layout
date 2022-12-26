//go:build wireinject
// +build wireinject

package api

import (
	"helloworld/dal/query"
	"helloworld/internal/conf"
	"helloworld/internal/controller"
	"helloworld/internal/pkg/store"
	"helloworld/internal/repository"
	"helloworld/internal/service"

	"github.com/google/wire"
	"gorm.io/gen"
)

func wireApp(*conf.Server, store.Config, ...gen.DOOption) (*APP, func(), error) {
	panic(wire.Build(
		store.NewMySQL,
		query.Use,
		repository.ProviderSet,
		service.ProviderSet,
		controller.ProviderSet,
		newApp,
	))
	return &APP{}, nil, nil
}
