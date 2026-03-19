package usecase

import (
	"errors"
	"hello-world/internal/entity"
)

var (
	ErrInvalidCEP   = errors.New("invalid zipcode")
	ErrCEPNotFound  = errors.New("can not find zipcode")
	ErrInternal     = errors.New("internal server error")
)

type LocationProvider interface {
	GetCityByCEP(cep string) (string, error)
}

type WeatherProvider interface {
	GetTemperature(city string) (float64, error)
}

type WeatherByCepUseCase struct {
	Location LocationProvider
	Weather  WeatherProvider
}

func NewWeatherByCepUseCase(l LocationProvider, w WeatherProvider) *WeatherByCepUseCase {
	return &WeatherByCepUseCase{Location: l, Weather: w}
}

func (uc *WeatherByCepUseCase) Execute(cep string) (*entity.Temperature, error) {
	if !entity.IsValidCEP(cep) {
		return nil, ErrInvalidCEP
	}

	city, err := uc.Location.GetCityByCEP(cep)
	if err != nil {
		if err.Error() == "cep not found" {
			return nil, ErrCEPNotFound
		}
		return nil, ErrInternal
	}

	tempC, err := uc.Weather.GetTemperature(city)
	if err != nil {
		return nil, ErrInternal
	}

	temp := entity.NewTemperature(tempC)
	return &temp, nil
}
