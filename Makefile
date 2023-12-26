MIGRATIONS_FOLDER = $(PWD)/db/migrations

.PHONY: test
test:
	go test -v ./...

.PHONY: build
build: 
	go build -o bin/restuarant

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
