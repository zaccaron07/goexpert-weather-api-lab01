package usecase_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zaccaron07/goexpert-weather-api-lab01/internal/entity"
	"github.com/zaccaron07/goexpert-weather-api-lab01/internal/usecase"
)

type MockWeatherRepository struct {
	mock.Mock
}

func (m *MockWeatherRepository) GetByCityName(cityName string) (entity.Weather, error) {
	args := m.Called(cityName)
	return args.Get(0).(entity.Weather), args.Error(1)
}

func TestGetWeatherUseCase(t *testing.T) {
	t.Run("should return weather data given a valid input", func(t *testing.T) {
		cityName := "Içara"

		mockWeatherRepository := new(MockWeatherRepository)

		mockWeatherRepository.On("GetByCityName", cityName).Return(entity.Weather{
			TemperatureCelsius:    25.0,
			TemperatureFahrenheit: 77.0,
			TemperatureKelvin:     298.15,
		}, nil)

		getWeatherInput := usecase.WeatherInput{
			CityName: cityName,
		}
		getWeatherUseCase := usecase.NewGetWeatherUseCase(mockWeatherRepository)
		getWeatherOutput, err := getWeatherUseCase.Execute(getWeatherInput)

		assert.NoError(t, err)
		assert.Equal(t, getWeatherOutput.TemperatureCelsius, 25.0)
		assert.Equal(t, getWeatherOutput.TemperatureFahrenheit, 77.0)
		assert.Equal(t, getWeatherOutput.TemperatureKelvin, 298.15)
		mockWeatherRepository.AssertExpectations(t)
	})

	t.Run("should return error when city not found", func(t *testing.T) {
		cityName := "Unknown City"

		mockWeatherRepository := new(MockWeatherRepository)

		mockWeatherRepository.On("GetByCityName", cityName).Return(entity.Weather{}, errors.New("city not found"))

		getWeatherInput := usecase.WeatherInput{
			CityName: cityName,
		}
		getWeatherUseCase := usecase.NewGetWeatherUseCase(mockWeatherRepository)
		getWeatherOutput, err := getWeatherUseCase.Execute(getWeatherInput)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "city not found")
		assert.Equal(t, getWeatherOutput, usecase.WeatherOutput{})
		mockWeatherRepository.AssertExpectations(t)
	})
}
