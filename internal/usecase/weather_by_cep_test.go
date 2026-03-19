package usecase

import (
	"errors"
	"testing"
)

type MockLocation struct {
	city string
	err  error
}

func (m *MockLocation) GetCityByCEP(cep string) (string, error) {
	return m.city, m.err
}

type MockWeather struct {
	temp float64
	err  error
}

func (m *MockWeather) GetTemperature(city string) (float64, error) {
	return m.temp, m.err
}

func TestWeatherByCepUseCase_Execute(t *testing.T) {
	t.Run("invalid cep", func(t *testing.T) {
		uc := NewWeatherByCepUseCase(&MockLocation{}, &MockWeather{})
		_, err := uc.Execute("123")
		if err != ErrInvalidCEP {
			t.Errorf("Expected ErrInvalidCEP, got %v", err)
		}
	})

	t.Run("cep not found", func(t *testing.T) {
		uc := NewWeatherByCepUseCase(&MockLocation{err: errors.New("cep not found")}, &MockWeather{})
		_, err := uc.Execute("11111111")
		if err != ErrCEPNotFound {
			t.Errorf("Expected ErrCEPNotFound, got %v", err)
		}
	})

	t.Run("success", func(t *testing.T) {
		uc := NewWeatherByCepUseCase(&MockLocation{city: "São Paulo"}, &MockWeather{temp: 28.5})
		temp, err := uc.Execute("01153000")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if temp.C != 28.5 {
			t.Errorf("Expected C to be 28.5, got %f", temp.C)
		}
	})
}
