# BookSPI

## Description

This is Books API with simple hash validation

## Architecture

![Clean Architecture](architecture.png)
(For this case is delivery only / cmd/web/...)
1. External system perform request (HTTP, gRPC, Messaging, etc)
2. The Delivery creates various Model from request data
3. The Delivery calls Service, and execute it using Model data
4. The Service create Entity data for the business logic
5. The Service calls Repository, and execute it using Entity data
6. The Repository use Entity data to perform database operation
7. The Repository perform database operation to the database
8. The Service create various Model for Gateway or from Entity data
9. The Service calls Gateway, and execute it using Model data
10. The Gateway using Model data to construct request to external system 
11. The Gateway perform request to external system (HTTP, gRPC, Messaging, etc)

## Tech Stack

- Golang : https://github.com/golang/go
- Postgres (Database) : https://github.com/go-gorm/postgres

## Framework & Library

- Gin (HTTP Framework) : https://github.com/gin-gonic/gin
- GRPC (Gateway between services and handler): https://google.golang.org/grpc
- GORM (ORM) : https://github.com/go-gorm/gorm
- Viper (Configuration) : https://github.com/spf13/viper
- Go Playground Validator (Validation) : https://github.com/go-playground/validator

## API Spec

All API Spec is in [docs folder](docs)

## Run Application

### Run web server

#### Copy env
```bash
cp .\.env.example .\.env
```
#### Running the server
```bash
go run .\cmd\web\user 
```

```bash
go run .\cmd\web\author
```

```bash
go run .\cmd\web\category
```

```bash
go run .\cmd\web\book 
```


### Docker compose

```bash
docker compose up -d
```

### Access Docs
Either use swagger in [docs folder](docs), or copy all collection and environment from [postman folder](postman)postman

### MVP
a) Book Management
- CRUD with connection to Author, Category services
b) Author Management
- CRUD with connection to Book services
c) Book Categories Management
- CRUD with connection to Book services
d) Book Stock Management
- Integrated with Book Management
e) Borrowing and Returning Books
- CRUD with connection to Author, Category services and Book Repository
f) Search and Recommendation feature