FROM golang:1.21

# Установим рабочую директорию в контейнере
WORKDIR /app

COPY . .

# Загружаем зависимости
RUN go mod download

# Копируем исходный код проекта в рабочую директорию

# Собираем приложение
RUN go build -o /awesomeproject

# Запускаем приложение
CMD ["/awesomeproject"]
