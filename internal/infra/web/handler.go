package web

import (
	"encoding/json"
	"hello-world/internal/usecase"
	"net/http"
	"strings"
)

type WeatherHandler struct {
	UseCase *usecase.WeatherByCepUseCase
}

func NewWeatherHandler(uc *usecase.WeatherByCepUseCase) *WeatherHandler {
	return &WeatherHandler{UseCase: uc}
}

func (h *WeatherHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		path := strings.TrimPrefix(r.URL.Path, "/")
		if path != "" {
			cep = path
		}
	}

	temp, err := h.UseCase.Execute(cep)
	if err != nil {
		switch err {
		case usecase.ErrInvalidCEP:
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		case usecase.ErrCEPNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(temp)
}
