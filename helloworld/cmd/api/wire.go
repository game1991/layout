//go:build wireinject
// +build wireinject

package api

import (
	"github.com/game1991/layout/helloworld/dal/query"
	"github.com/game1991/layout/helloworld/internal/conf"
	"github.com/game1991/layout/helloworld/internal/controller"
	"github.com/game1991/layout/helloworld/internal/pkg/store"
	"github.com/game1991/layout/helloworld/internal/repository"
	"github.com/game1991/layout/helloworld/internal/service"

	"github.com/google/wire"
	"gorm.io/gen"
)

func wireApp(*conf.Server, *conf.Session, store.Config, ...gen.DOOption) (*APP, func(), error) {
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
