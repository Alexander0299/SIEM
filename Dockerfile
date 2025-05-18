# Используем официальный образ Golang
FROM golang:1.23.4

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum и скачиваем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# Собираем проект
RUN go build -o siem .

# Устанавливаем переменные окружения
ENV CGO_ENABLED=0

# Запускаем приложение
CMD ["./siem"]
