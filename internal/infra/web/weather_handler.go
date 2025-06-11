package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zaccaron07/goexpert-weather-api-lab01/internal/entity"
	"github.com/zaccaron07/goexpert-weather-api-lab01/internal/infra/repo"
	"github.com/zaccaron07/goexpert-weather-api-lab01/internal/usecase"
)

type WebWeatherHandler struct {
	ZipcodeRepository entity.ZipcodeRepositoryInterface
}

func NewWebWeatherHandler() *WebWeatherHandler {
	return &WebWeatherHandler{
		ZipcodeRepository: repo.NewZipcodeRepository(),
	}
}

func (h *WebWeatherHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	zipcode := chi.URLParam(r, "zipcode")
	fmt.Println("Fetching weather for Zipcode:", zipcode)

	getZipcodeInput := usecase.ZipcodeInput{
		CEP: zipcode,
	}

	getZipcodeUseCase := usecase.NewGetZipcodeUseCase(h.ZipcodeRepository)
	getZipcodeOutput, err := getZipcodeUseCase.Execute(getZipcodeInput)

	if err != nil {
		if err.Error() == "invalid zipcode" {
			http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
			return
		}

		http.Error(w, fmt.Sprintf("error fetching location: %v", err), http.StatusInternalServerError)
		return
	}
	if getZipcodeOutput.Localidade == "" {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getZipcodeOutput)
}
