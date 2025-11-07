# Documentation for **Identity** microservice

Cервис для управления пользователями и их внешними идентификаторами. Используется как центр идентификации: регистрации/привязки Telegram и чтения профиля по внутреннему user_id или telegram_id.

# Structure 

Для построения структуры сервиса была использована **Чистая архитектура**

```
cmd/                    — точка входа
internal/
  app/                  — запуск сервиса и создание всех компонентов
  domain/               — модели, инварианты
  usecase/              — сценарии, интерфейсы внешних адаптеров
  adapters/
    grpc/               — сервер, валидаторы, мапперы
    http/               — /metrics
    txmanager/          — менеджер транзакций
  repository/           — репозиторий PostgreSQL
proto/identity/         — .proto + buf
proto-codegen/          — сгенерированный gRPC код
```
