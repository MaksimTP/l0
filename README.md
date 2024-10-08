# Демонстрационный сервис для отображения данных о заказе

## Описание

Данный сервис предназначен для демонстрации работы с заказами, получаемыми из Kafka. Он включает в себя интерфейс для отображения информации о заказах и использует PostgreSQL для хранения данных. Сервис реализует кэширование, что позволяет быстро получать информацию без постоянных запросов к базе данных.

### Запуск

- Клонирование репозитория

```bash
git clone https://github.com/MaksimTP/l0.git && cd l0
```

#### Вариант №1

- Через docker-compose (у меня есть проблемы с ним, в контейнерах network issues):

```bash
docker-compose up --build order_service
```
#### Вариант №2

- Поднимать контейнеры локально. Начнем с брокера сообщений.

```bash
docker run --name=broker -p 9092:9092  apache/kafka:latest
```

- БД.

```bash
docker run --name=wb -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -e POSTGRES_DB=wb_lvl0 -v $(pwd)/init/init.sql:/docker-entrypoint-initdb.d/init.sql -d postgres:16
```

- Запуск сервиса

```bash
go run cmd/main.go
```

- Запуск продюсера (один запуск - одно сообщение в топик)

```bash
go run internal/kafka/producer/producer.go model2.json
```

- Работа http-серверa. Зайти в браузер, ввести адресной строке: `http://localhost:8000/`.  Сервис будет доступен по адресу http://localhost:5000/orders/{id}. Форма выполняет GET-запрос.


### Генерация фейковых данных

```bash
go run internal/data_generation.go
```

### Тесты

```bash
go test -v ./...
```

### Линтер

```bash
go vet ./...
```