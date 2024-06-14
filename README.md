# Go Fiber Postgres Boilerplate
This is an example of Go project with Fiber and Postgres

## Technology used

- Go v1.21
- PostgreSQL v15

## External Libraries

- Gorm
- Fiber
- Manual Dependency injection -> check (internal/infrastructure/injector.go)
- [Zap](https://github.com/uber-go/zap) - Logger (SugaredLogger)
- [golang-migrate](https://github.com/golang-migrate/migrate) - DB migrations
- Other libraries

## Folder Structure

- `cmd`: main file to call the app
- `configs`: helper to load .env & all the configuration needed
- `pkg`: shared packages
- `database/migrations`: all database migration files created by [golang-migrate](https://github.com/golang-migrate/migrate)
- `internal`: internal modules
  - `infrastructure`: main app to initiate server & injecting dependency
  - `dto`: data to object usually containing Request Response structs or any other struct convertions1
  - `repository`: repository - anything with db interaction
  - `usecase`: usecase - all business logic should be here so it will bridge between repository - usecase - handler/api handler
  - `api`: api handler - api routes & handler definition


## Local Development

### Prerequisite
1. Postgres/Postgres-Docker installed
2. Go v1.21 installed
3. golang-migrate installed [golang-migrate](https://github.com/golang-migrate/migrate)

### Setup
1. Clone this repository
2. Copy `.env.example` to `.env`
3. Run `go mod tidy` to download all dependencies or Run `go mod vendor` to also lock version dependencies so all package will installed in this project folder rather than on GO ENV.
4. if your running locally don't forget to migrate database first with `make migrate`, but for non local (based on .env APP_ENV) it will AutoMigrate
5. Try to run the app `make run`
6. or if you want to use live reload you can use golang air and just type `air` or using nodemon and type `make live`

### Dependency Injection
now still manual dependencies injection check -> internal/infrastructure/injector.go
