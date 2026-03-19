# Sistema de Clima por CEP (Go Expert)

Este é um sistema desenvolvido em Go que recebe um CEP, identifica a cidade utilizando a [ViaCEP](https://viacep.com.br/) e retorna a temperatura atual em graus Celsius, Fahrenheit e Kelvin através da [WeatherAPI](https://www.weatherapi.com/).

## Como rodar localmente com Docker

### Pré-requisitos
- Docker instalado.
- Chave de API da [WeatherAPI](https://www.weatherapi.com/).

### Executando a Aplicação
1. Construa a imagem do Docker:
   ```bash
   docker build -t go-weather-cep .
   ```

2. Execute o container, não se esquecendo de preencher sua API Key no lugar de `<SUA_API_KEY>`:
   ```bash
   docker run -p 8080:8080 -e WEATHER_API_KEY="<SUA_API_KEY>" go-weather-cep
   ```

3. Teste o endpoint informando o CEP na URL ou por Query String:
   ```bash
   curl http://localhost:8080/01153000
   # ou
   curl "http://localhost:8080/?cep=01153000"
   ```

## Rodando os Testes Automatizados

O projeto utiliza Clean Architecture e contém testes unitários focados nas conversões das entidades e fluxos do use case.
Para executar todos os testes, basta rodar:

```bash
go test ./...
```

## Acesso ao Cloud Run

O projeto está pronto para ser publicado no Google Cloud Run através do `Dockerfile` multi-stage na raiz do repositório.

A URL pública da aplicação publicada no Google Cloud Run encontra-se abaixo:
https://cloudrun-goexpert-7pgcz3avoa-uc.a.run.app/

Exemplo de run com CEP valido:
https://cloudrun-goexpert-7pgcz3avoa-uc.a.run.app/99709292