DB_HOST=db
DB_PORT=5432
DB_USER=gwp
DB_PASSWORD=gwp
DB_NAME=gwp
DB_CONN=postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable

.PHONY: run
run:
	docker-compose up --build -d

.PHONY: start
start:
	docker-compose exec app realize start --run

# migrationファイルの作成
.PHONY: migrate-create
migrate-create:
	docker-compose exec app migrate create -ext sql -dir migrations ${FILENAME}

# migrationの実行
.PHONY: migrate-up
migrate-up:
	docker-compose exec app migrate --source file://migrations --database ${DB_CONN} up

# migration(rollback)の実行
.PHONY: migrate-down
migrate-down:
	docker-compose exec app migrate --source file://migrations --database ${DB_CONN} down 1


# go generate
.PHONY: generate
generate:
	docker-compose exec app go run github.com/99designs/gqlgen generate

.PHONY: tidy
tidy:
	docker-compose exec app go mod tidy

.PHONY: getgql
getgql:
	docker-compose exec app go get github.com/99designs/gqlgen@v0.17.2

.PHONY: goinstall
goinstall:
	docker-compose exec app go install