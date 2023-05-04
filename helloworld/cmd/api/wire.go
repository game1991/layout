//go:build wireinject
// +build wireinject

package api

import (
	"git.xq5.com/golang/helloworld/dal/query"
	"git.xq5.com/golang/helloworld/internal/conf"
	"git.xq5.com/golang/helloworld/internal/controller"
	"git.xq5.com/golang/helloworld/internal/pkg/store"
	"git.xq5.com/golang/helloworld/internal/repository"
	"git.xq5.com/golang/helloworld/internal/service"

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
