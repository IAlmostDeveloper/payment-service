# payment-service

JSON API, имитирующий работу платежной системы

# Зависимости для сборки
- github.com/gorilla/mux
- github.com/go-openapi/runtime/middleware
- github.com/mattn/go-sqlite3
- github.com/google/uuid
# Инструкции по сборке
## Go build
```go build main.go && main.exe```
## Docker
```docker run -it -p 8080:8080 ialmostdeveloper/payment-service```
## Docker-compose
```docker-compose up```
