package viacep

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ViaCEPResponse struct {
	Localidade string `json:"localidade"`
	Erro       bool   `json:"erro"`
}

type ViaCEPClient struct{}

func NewViaCEPClient() *ViaCEPClient {
	return &ViaCEPClient{}
}

func (c *ViaCEPClient) GetCityByCEP(cep string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("viacep error: status %d", resp.StatusCode)
	}

	var data ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	if data.Erro {
		return "", errors.New("cep not found")
	}

	return data.Localidade, nil
}
