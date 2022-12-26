package controller

import "github.com/google/wire"

// ProviderSet is repository providers.
var ProviderSet = wire.NewSet(
	NewHandler,
)
