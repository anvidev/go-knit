include .env

export GOOSE_DRIVER = ${DATABASE_DRIVER}
export GOOSE_DBSTRING = ${DATABASE_URL}
export GOOSE_MIGRATION_DIR = ${DATABASE_MIGRATION_DIR}

build:
	tailwindcss -i static/css/styles.css -o static/styles.css
	@templ generate
	@go build -o tmp/main.go ./cmd/app/main.go

test:
	@go test -v ./...
	
run: build
	@./tmp/main.go

css:
	@tailwindcss -i internal/view/css/styles.css -o static/css/styles.css --watch

templ:
	@templ generate --watch --proxy http://localhost:3000

db-create:
	@goose create ${NAME} sql

db-status:
	@goose status

db-up:
	@goose up

db-down:
	@goose down

db-validate:
	@goose validate

