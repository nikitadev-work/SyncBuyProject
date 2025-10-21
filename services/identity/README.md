# Documentation for **Identity** microservice

# Structure 

Для построения структуры сервиса была использована **Чистая архитектура**

```
/identity
├── cmd                 
│   └── main.go         // точка старта конфигурации сервиса
├── internal            
│   ├── app             
│   │   └── app.go      // точка запуска сервиса + graceful shutdown
│   ├── interfaces      // http, grpc, repo, etc.
│   ├── domain          // основная бизнес логика
│   └── usecase
├── go.mod
├── README.md
└── tests
```