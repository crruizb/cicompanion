DB_URL=postgresql://user:user@localhost:5432/cicompanion?sslmode=disable

run:
	go run main.go

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

test:
	go test -v -cover -short ./...