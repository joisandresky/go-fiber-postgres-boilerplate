postgresconn=postgres://<USER>:<PASSWORD>@<HOST>:<PORT>/<DB_NAME>?sslmode=disable

.PHONY: migration-status migrate unmigrate migration migrate-force live run

run:
	go run ./cmd/main.go

# live command is using Nodemon, so make sure you have nodemon installed
# either way you can use golang air, and just run with "air"
live:
	nodemon --delay 1 --exec go run ./cmd/main.go --signal SIGTERM

migration-status:
	migrate -database $(postgresconn) -path ./database/migrations version

migrate:
	migrate -database $(postgresconn) -path ./database/migrations up

unmigrate:
	migrate -database $(postgresconn) -path ./database/migrations down 1

migrate-force:
	@read -p "Enter Migration Version: (ex: 20230901112355) " name; \
	migrate -database $(postgresconn) -path ./database/migrations force $$name

migration:
	@read -p "Enter Migration Name: (with underscore) " name; \
		migrate create -ext sql -dir ./database/migrations $$name