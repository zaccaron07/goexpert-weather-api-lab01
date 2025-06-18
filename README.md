# GoExpert Weather API Lab 01

Uma API REST em Go que recebe um CEP brasileiro, identifica a cidade e retorna o clima atual (temperatura em Celsius, Fahrenheit e Kelvin). Construída com clean architecture e pronta para deploy no Google Cloud Run.

## Funcionalidades
- Recebe um CEP válido de 8 dígitos
- Busca a cidade usando a API ViaCEP
- Consulta o clima atual usando a WeatherAPI
- Retorna temperatura em Celsius, Fahrenheit e Kelvin
- Tratamento de erros para CEP inválido ou não encontrado
- Compatível com Docker e Podman
- Pronto para Cloud Run

## Requisitos
- Go 1.22+
- Chave da [WeatherAPI](https://www.weatherapi.com/)
- Docker ou Podman (para conteinerização)

## Variáveis de Ambiente
- `WEATHER_API_KEY`: Sua chave da WeatherAPI (obrigatória)

## Executando Localmente

1. **Defina a variável de ambiente:**
   - No PowerShell:
     ```powershell
     $env:WEATHER_API_KEY="sua_chave_aqui"; go run ./cmd/lab01
     ```
   - No Linux/macOS (bash/zsh):
     ```sh
     WEATHER_API_KEY=sua_chave_aqui go run ./cmd/lab01
     ```
   - Ou crie um arquivo `.env` (apenas para desenvolvimento local):
     ```env
     WEATHER_API_KEY=sua_chave_aqui
     ```
     E use [github.com/joho/godotenv](https://github.com/joho/godotenv) se quiser carregá-la automaticamente.

2. **Execute os testes com cobertura:**
   ```sh
   go test ./... -coverprofile=coverage.out
   go tool cover -func=coverage
   go tool cover -html=coverage
   ```

## Build e Execução com Docker/Podman

1. **Build da imagem:**
   ```sh
   podman build -t goexpert-weather-api .
   # ou
   docker build -t goexpert-weather-api .
   ```

2. **Execute o container:**
   ```sh
   podman run --rm -p 8080:8080 -e WEATHER_API_KEY=sua_chave_aqui goexpert-weather-api
   # ou
   docker run --rm -p 8080:8080 -e WEATHER_API_KEY=sua_chave_aqui goexpert-weather-api
   ```

## Uso da API

- **Endpoint:** `GET /zipcode/{zipcode}/weather`
- **Resposta de Sucesso:**
  - HTTP 200
  - Body: `{ "temp_C": 28.5, "temp_F": 83.3, "temp_K": 301.65 }`
- **CEP Inválido:**
  - HTTP 422
  - Body: `invalid zipcode`
- **CEP Não Encontrado:**
  - HTTP 404
  - Body: `can not find zipcode`

## Estrutura do Projeto

- `cmd/lab01/` — Ponto de entrada principal
- `internal/entity/` — Entidades de domínio e interfaces
- `internal/infra/repo/` — Repositórios de APIs externas
- `internal/infra/web/` — Handlers HTTP e servidor web
- `internal/usecase/` — Casos de uso da aplicação

## Link da entrega
https://goexpert-weather-api-lab-01-516049214932.us-central1.run.app/zipcode/88820000/weather
