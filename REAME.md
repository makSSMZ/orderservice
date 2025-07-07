# OrderService
**OrderService** — учебный микросервис на Go, предназначенный для приёма заказов из Kafka, сохранения их в PostgreSQL, кэширования в памяти и предоставления HTTP API для получения информации о заказах.

Этот проект создан в рамках обучения разработке микросервисов на Go.

## Быстрый старт
### 1. Клонируй репозиторий
```bash
git clone git@github.com:makSSMZ/orderservice.git
cd orderservice
```

### 2. Запусти через Docker Compose
```bash
docker-compose up --build
```
#### Это поднимет:
- Kafka + Zookeeper
- Kafka UI 
- PostgreSQL
- OrderService

### API
```
GET /orders/{order_uid} — Получить заказ по ID
```

#### Пример:
```bash
curl http://localhost:8081/order/test123
```
