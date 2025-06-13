package entity_test

import (
	"testing"

	"github.com/zaccaron07/goexpert-weather-api-lab01/internal/entity"
)

func TestWeatherEntity(t *testing.T) {
	t.Run("should create a valid Weather entity", func(t *testing.T) {
		temperatureCelsius := 25.0

		w := entity.NewWeather(temperatureCelsius)
		if w.TemperatureCelsius != temperatureCelsius {
			t.Errorf("expected TemperatureCelsius %f, got %f", temperatureCelsius, w.TemperatureCelsius)
		}
		if w.TemperatureFahrenheit != (temperatureCelsius*1.8)+32 {
			t.Errorf("expected TemperatureFahrenheit %f, got %f", (temperatureCelsius*1.8)+32, w.TemperatureFahrenheit)
		}
		if w.TemperatureKelvin != temperatureCelsius+273 {
			t.Errorf("expected TemperatureKelvin %f, got %f", temperatureCelsius+273, w.TemperatureKelvin)
		}
	})
}
