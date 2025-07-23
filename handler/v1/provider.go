package v1

import (
	"github.com/google/wire"
)

type APIHandlers struct {
	AuthHandler *AuthHandler
}

var ProviderSet = wire.NewSet(
	NewAuthHandler,
	wire.Struct(new(APIHandlers), "*"),
)
