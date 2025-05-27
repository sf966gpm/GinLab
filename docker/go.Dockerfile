# Устанавливаем зависимости и новую версию Air
FROM golang:1.24.3-alpine AS dev

# Устанавливаем Air и зависимости
RUN apk add --no-cache git curl \
    && go install github.com/air-verse/air@latest

WORKDIR /app

COPY ./app .

# Устанавливаем зависимости
RUN go mod download

CMD ["air", "-c", ".air.toml"]