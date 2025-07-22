package repo

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewUserRepo, NewAuthRepo, NewRouteRepo, NewUserTripRepo, NewMomentRepo, NewFollowRepo, NewUserLocationRepo)
