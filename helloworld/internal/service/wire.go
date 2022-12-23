package service

import "github.com/google/wire"

// ProviderSet .
var ProviderSet = wire.NewSet(
	NewUserSrv,
	NewGreeterSrv,
)
