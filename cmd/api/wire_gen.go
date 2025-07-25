// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"framework/apm"
	"framework/config"
	"framework/db/pg"
	v1 "framework/handler/v1"
	"framework/log"
	"framework/repo"
	"framework/server/http"
	"framework/usecase"
)

// Injectors from wire.go:

func createApp() (*App, error) {
	configConfig, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	logger := log.NewLogger(configConfig)
	echo := http.NewEcho(logger, configConfig)
	httpServer := &http.HTTPServer{
		Echo: echo,
	}
	tracer, err := apm.NewTracer(configConfig)
	if err != nil {
		return nil, err
	}
	db, err := pg.NewDB()
	if err != nil {
		return nil, err
	}
	gormDB := pg.ProvideGormDB(db)
	authRepo := repo.NewAuthRepo(gormDB)
	authUsecase := usecase.NewAuthUsecase(authRepo)
	authHandler := v1.NewAuthHandler(echo, authUsecase)
	apiHandlers := &v1.APIHandlers{
		AuthHandler: authHandler,
	}

	app := &App{
		HTTPServer: httpServer,
		Config:     configConfig,
		Tracer:     tracer,
		Handlers:   apiHandlers,
	}
	return app, nil
}

// wire.go:

type App struct {
	HTTPServer *http.HTTPServer
	Config     *config.Config
	Tracer     *apm.Tracer
	Handlers   *v1.APIHandlers
}
