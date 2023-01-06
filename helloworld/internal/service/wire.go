package service

import "github.com/google/wire"

// ProviderSet .
var ProviderSet = wire.NewSet(
	NewService,
	NewUserSrv,
	NewGreeterSrv,
)
