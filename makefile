get-docs:
	go get -u github.com/swaggo/swag/cmd/swag

docs: get-docs
	swag init --dir api --parseDependency --output docs

build:
	go build -o bin/restapi cmd/main.go

run:
	go run cmd/main.go

test:
	go test -v ./test/...

build-docker: build
	docker build . -t api-rest

run-docker: build-docker
	docker run -p 3000:3000 api-rest

gen-models:
	jet -source=postgres -dsn="user= password= host= port=4000 dbname=sslmode=disable" -schema=public -path=./.gen

migrate-up:
	migrate -path internal/pkg/db/migrations -database "postgres://:4000/?sslmode=disable" up

migrate-down:
	migrate -path internal/pkg/db/migrations -database "postgres://:4000/?sslmode=disable" down

migrate-create:
	migrate create -ext sql -dir internal/pkg/db/migrations -seq $(name)
