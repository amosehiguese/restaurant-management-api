MIGRATIONS_FOLDER = $(PWD)/db/migrations
SWAG_PATH ?= $(go env GOPATH)
VERSION ?= $(shell git describe --match 'v[0-9]*' --tags --always)
DBUSER = $(DB_USER)
DBPORT = $(DB_PORT)
DBNAME = $(DB_NAME)
DBPASSWORD = $(DB_PASSWORD)


.PHONY: test
test:
	go test -v ./...

.PHONY: build
build: 
	@go build -ldflags "-X main.version=$(VERSION)" -o bin/restuarant

.PHONY:start
start: build
	./bin/restuarant

.PHONY:mig-create
mig-create:
	migrate create -ext sql -dir migrations -seq create_restaurant_tables

.PHONY: mig-up
mig-up:
	migrate -database "$(DATABASE_URL)" -path $(MIGRATIONS_FOLDER) up

.PHONY: mig-down
mig-down:
	migrate -database "$(DATABASE_URL)"  -path $(MIGRATIONS_FOLDER) down

.PHONY: mig-force
mig-force:
	migrate -database "$(DATABASE_URL)" -path $(MIGRATIONS_FOLDER) force 1

docker.postgres:
	docker run --rm -d \
		--name postgresql \
		-e POSTGRES_USER=$(DBUSER) \
		-e POSTGRES_PASSWORD=$(DBPASSWORD) \
		-e POSTGRES_DB=$(DBNAME) \
		-p 5432:5432 \
		bitnami/postgresql:latest

docker.redis:
	docker run --rm -d \
		--name myredis \
		--network dev-network \
		-p 6379:6379 \
		redis


generate-docs:
	go run "$(SWAG_PATH)/bin/swag" init 