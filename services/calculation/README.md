# Documentation for **Calculation** microservice

# Structure 

Для построения структуры сервиса была использована **Чистая архитектура**

```
/calculation
├── cmd                 
│   └── main.go         // точка старта конфигурации сервиса
├── internal            
│   ├── app             
│   │   └── app.go      // точка запуска сервиса + graceful shutdown
│   ├── adapters        // http, grpc, repo, etc.
│   │   ├── http
│   │   ├── repo
│   │   └── grpc
│   ├── domain          // основная бизнес логика
│   ├── infra           // logger, metrics, etc.
│   │   ├── config.go   
│   │   └── logger.go
│   │   
│   └── usecase
├── go.mod
├── README.md
└── tests
```