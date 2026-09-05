include .env
export

.PHONEY: run stop

run: db-up
	go run ./cmd/api

stop: db-down
	echo "Stop Application"

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	echo "TBD"

db-up:
	echo "Start Local Database"
	docker compose up -d

db-down:
	echo "Stop Local Database"
	docker compose down

migrate-up:
	migrate -database $(ALPHA_DATABASE_URL) -path db/migrations up

migrate-down:
	migrate -database ${ALPHA_DATABASE_URL} -path db/migrations down

migrate-version:
	migrate -version ${ALPHA_DATABASE_URL}

sqlc:
	echo "TBD"
