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
	echo "TBD"

migrate-down:
	echo "TBD"

sqlc:
	echo "TBD"
