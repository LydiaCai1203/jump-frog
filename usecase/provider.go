package usecase

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewAuthUsecase, NewUserUsecase, NewRouteUsecase, NewUserTripUsecase, NewMomentUsecase, NewFollowUsecase, NewUserLocationUsecase)
