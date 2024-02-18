MIGRATIONS_FOLDER = $(PWD)/db/migrations
ENVIRONMENT = $(ENVIRONMENT)
VERSION ?= $(shell git describe --match 'v[0-9]*' --tags --always)
TAG ?= 1.0.0



.PHONY: test
test:
	go test -v ./...

.PHONY: build
build: 
	go build -ldflags "-X main.version=$(VERSION)" -o bin/restuarant

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

build-docker:
	docker build -t github.com/amosehiguese/restaurant-api:$(TAG) .

helm-dep:
	helm dependency update

helm-install: 
	sops -d deploy/restaurant/${ENVIRONMENT}-secrets.yaml > temp-${ENVIRONMENT}-secrets.yaml && helm -n ${ENVIRONMENT} upgrade --install ${ENVIRONMENT}-restaurant-api --values deploy/restaurant/values.yaml --values temp-${ENVIRONMENT}-secrets.yaml ./deploy/restaurant && rm temp-${ENVIRONMENT}-secrets.yaml

docker.redis:
	docker run --rm -d \
		--name myredis \
		--network dev-network \
		-p 6379:6379 \
		redis

docker.stop.postgres:
	docker stop postgresql

docker.stop.redis:
	docker stop myredis

swag:
	swag init

