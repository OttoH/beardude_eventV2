init: deps
dev: dev
db: mongo
build: build

.PHONY: deps
deps:
	brew install glide #curl https://glide.sh/get | sh
	glide install

.PHONY: dev
dev:
	go run main.go routers.go authMiddleware.go

.PHONY: build
build:
	rm -rf ./bin/qpetEngine
	go build -o ./bin/qpetEngine main.go routers.go authMiddleware.go
	GIN_MODE=release
	./bin/qpetEngine

.PHONY: mongo
mongo:
		mongod --dbpath data/db
