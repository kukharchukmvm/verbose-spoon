# stage 1: build
FROM golang:1.21-alpine AS build
WORKDIR /app

# Копіюємо main.go з кореня проекту
COPY ../main.go .

# Створюємо бінарник
RUN go build -o main main.go

# stage 2: run
FROM alpine:latest
WORKDIR /app

# Копіюємо бінарник з етапу збірки
COPY --from=build /app/main .

# Команда запуску
CMD ["./main"]
