//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/zaccaron07/goexpert-weather-api-lab01/internal/entity"
	"github.com/zaccaron07/goexpert-weather-api-lab01/internal/infra/repo"
	"github.com/zaccaron07/goexpert-weather-api-lab01/internal/infra/web"
)

var setZipcodeRepository = wire.NewSet(
	repo.NewZipcodeRepository,
	wire.Bind(new(entity.ZipcodeRepositoryInterface), new(repo.ZipcodeRepository)),
)

var setWeatherRepository = wire.NewSet(
	repo.NewWeatherRepository,
	wire.Bind(new(entity.WeatherRepositoryInterface), new(repo.WeatherRepository)),
)

func NewWeatherHandler() *web.WebWeatherHandler {
	wire.Build(web.NewWebWeatherHandler)
	return &web.WebWeatherHandler{}
}
