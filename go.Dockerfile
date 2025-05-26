# Устанавливаем зависимости и новую версию Air
FROM golang:1.24.3-alpine AS dev

# Устанавливаем Air и зависимости
RUN apk add --no-cache git curl \
    && go install github.com/air-verse/air@latest

WORKDIR /app

# Копируем всё кроме исключённых в .dockerignore
COPY . .

# Устанавливаем зависимости
RUN go mod download

CMD ["air", "-c", ".air.toml"]

# === Финальный образ для продакшена ===
FROM golang:1.24.3-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/ginlab .

FROM alpine:3.18 AS prod
COPY --from=builder /app/ginlab /app/ginlab
EXPOSE 8080
CMD ["/app/ginlab"]