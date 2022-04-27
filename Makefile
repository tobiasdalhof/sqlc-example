DB_URL=postgresql://default:secret@postgres:5432/default?sslmode=disable

migration:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrate_up:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

migrate_down:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

test:
	go test ./...

.PHONY: migration migrate_up migrate_down sqlc test
