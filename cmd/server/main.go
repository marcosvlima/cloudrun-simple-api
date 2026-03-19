package main

import (
	"fmt"
	"hello-world/internal/infra/viacep"
	"hello-world/internal/infra/weatherapi"
	"hello-world/internal/infra/web"
	"hello-world/internal/usecase"
	"log"
	"net/http"
	"os"
)

func main() {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		log.Println("WARNING: WEATHER_API_KEY is not set")
	}

	locationClient := viacep.NewViaCEPClient()
	weatherClient := weatherapi.NewWeatherAPIClient(apiKey)
	uc := usecase.NewWeatherByCepUseCase(locationClient, weatherClient)
	handler := web.NewWeatherHandler(uc)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.Handle("/", handler)

	fmt.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Server failed: %s\n", err)
	}
}
