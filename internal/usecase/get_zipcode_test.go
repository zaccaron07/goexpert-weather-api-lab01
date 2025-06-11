package usecase_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zaccaron07/goexpert-weather-api-lab01/internal/entity"
	"github.com/zaccaron07/goexpert-weather-api-lab01/internal/usecase"
)

type MockZipcodeRepository struct {
	mock.Mock
}

func (m *MockZipcodeRepository) Get(zipcode string) (entity.Zipcode, error) {
	args := m.Called(zipcode)
	return args.Get(0).(entity.Zipcode), args.Error(1)
}

func TestGetZipcodeUseCase(t *testing.T) {
	t.Run("should return zipcode data given a valid input", func(t *testing.T) {
		zipcode := "88820000"

		mockZipcodeRepository := new(MockZipcodeRepository)

		mockZipcodeRepository.On("Get", zipcode).Return(entity.Zipcode{
			CEP:        "88820-000",
			Bairro:     "",
			Localidade: "Içara",
			UF:         "SC",
		}, nil)

		getZipcodeInput := usecase.ZipcodeInput{
			CEP: zipcode,
		}
		getZipcodeUseCase := usecase.NewGetZipcodeUseCase(mockZipcodeRepository)
		getZipcodeOutput, err := getZipcodeUseCase.Execute(getZipcodeInput)

		assert.NoError(t, err)
		assert.Equal(t, getZipcodeOutput.CEP, "88820-000")
		assert.Equal(t, getZipcodeOutput.Bairro, "")
		assert.Equal(t, getZipcodeOutput.Localidade, "Içara")
		assert.Equal(t, getZipcodeOutput.UF, "SC")
		mockZipcodeRepository.AssertExpectations(t)
	})

	t.Run("should return invalid zipcode error", func(t *testing.T) {
		zipcode := "879164"

		mockZipcodeRepository := new(MockZipcodeRepository)

		getZipcodeInput := usecase.ZipcodeInput{
			CEP: zipcode,
		}
		getZipcodeUseCase := usecase.NewGetZipcodeUseCase(mockZipcodeRepository)
		getZipcodeOutput, err := getZipcodeUseCase.Execute(getZipcodeInput)

		assert.Error(t, err)
		assert.Equal(t, getZipcodeOutput.CEP, "")
		assert.Equal(t, getZipcodeOutput.Bairro, "")
		assert.Equal(t, getZipcodeOutput.Localidade, "")
		assert.Equal(t, getZipcodeOutput.UF, "")
		mockZipcodeRepository.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		zipcode := "88820000"

		mockZipcodeRepository := new(MockZipcodeRepository)
		mockZipcodeRepository.On("Get", zipcode).Return(entity.Zipcode{}, errors.New("repository error"))

		getZipcodeInput := usecase.ZipcodeInput{
			CEP: zipcode,
		}
		getZipcodeUseCase := usecase.NewGetZipcodeUseCase(mockZipcodeRepository)
		getZipcodeOutput, err := getZipcodeUseCase.Execute(getZipcodeInput)

		assert.Error(t, err)
		assert.Equal(t, getZipcodeOutput.CEP, "")
		assert.Equal(t, getZipcodeOutput.Bairro, "")
		assert.Equal(t, getZipcodeOutput.Localidade, "")
		assert.Equal(t, getZipcodeOutput.UF, "")
		mockZipcodeRepository.AssertExpectations(t)
	})
}
